package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"tcc-itpln/backend/internal/domain"
)

type transaksiRepo struct{ db *pgxpool.Pool }

func NewTransaksiRepository(db *pgxpool.Pool) TransaksiRepository { return &transaksiRepo{db} }

const transaksiCols = `id, user_id, kelas_id, pendaftaran_id, midtrans_order_id, COALESCE(midtrans_txn_id,''), jumlah::float8, status, COALESCE(metode_pembayaran,''), created_at, updated_at`

func scanTransaksi(row pgx.Row) (domain.Transaksi, error) {
	var t domain.Transaksi
	err := row.Scan(&t.ID, &t.UserID, &t.KelasID, &t.PendaftaranID, &t.MidtransOrderID, &t.MidtransTxnID, &t.Jumlah, &t.Status, &t.MetodePembayaran, &t.CreatedAt, &t.UpdatedAt)
	return t, err
}

func (r *transaksiRepo) Create(ctx context.Context, userID, kelasID, orderID string, jumlah float64) (string, error) {
	var id string
	err := r.db.QueryRow(ctx,
		`INSERT INTO transaksi (user_id, kelas_id, midtrans_order_id, jumlah) VALUES ($1,$2,$3,$4) RETURNING id`,
		userID, kelasID, orderID, jumlah).Scan(&id)
	return id, err
}

func (r *transaksiRepo) GetByOrderID(ctx context.Context, orderID string) (domain.Transaksi, error) {
	t, err := scanTransaksi(r.db.QueryRow(ctx, `SELECT `+transaksiCols+` FROM transaksi WHERE midtrans_order_id=$1`, orderID))
	if errors.Is(err, pgx.ErrNoRows) {
		return t, ErrNotFound
	}
	return t, err
}

func (r *transaksiRepo) ListByUser(ctx context.Context, userID string) ([]domain.Transaksi, error) {
	return r.query(ctx, `WHERE user_id=$1`, userID)
}

func (r *transaksiRepo) ListAll(ctx context.Context) ([]domain.Transaksi, error) {
	return r.query(ctx, ``)
}

func (r *transaksiRepo) query(ctx context.Context, where string, args ...any) ([]domain.Transaksi, error) {
	rows, err := r.db.Query(ctx, `SELECT `+transaksiCols+` FROM transaksi `+where+` ORDER BY created_at DESC`, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []domain.Transaksi{}
	for rows.Next() {
		t, err := scanTransaksi(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

func (r *transaksiRepo) Settle(ctx context.Context, orderID, txnID, metode string) (SettleResult, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return SettleResult{}, err
	}
	defer tx.Rollback(ctx)

	var status, userID, kelasID string
	err = tx.QueryRow(ctx, `SELECT status, user_id, kelas_id FROM transaksi WHERE midtrans_order_id=$1 FOR UPDATE`, orderID).
		Scan(&status, &userID, &kelasID)
	if errors.Is(err, pgx.ErrNoRows) {
		return SettleResult{}, ErrNotFound
	}
	if err != nil {
		return SettleResult{}, err
	}
	if status == "sukses" {
		return SettleResult{AlreadyDone: true}, nil
	}

	tag, err := tx.Exec(ctx, `
		UPDATE kelas SET peserta_terdaftar = peserta_terdaftar + 1,
		       status = CASE WHEN kuota > 0 AND peserta_terdaftar + 1 >= kuota THEN 'penuh' ELSE status END
		WHERE id = $1 AND status = 'aktif' AND (kuota = 0 OR peserta_terdaftar < kuota)`, kelasID)
	if err != nil {
		return SettleResult{}, err
	}
	if tag.RowsAffected() == 0 {
		return SettleResult{KuotaFull: true}, nil
	}

	var pid string
	err = tx.QueryRow(ctx, `INSERT INTO pendaftaran (user_id, kelas_id, status) VALUES ($1,$2,'aktif') RETURNING id`, userID, kelasID).Scan(&pid)
	if err != nil {
		var pgErr *pgconn.PgError
		if !errors.As(err, &pgErr) || pgErr.Code != "23505" {
			return SettleResult{}, err
		}
		if _, err := tx.Exec(ctx, `
			UPDATE kelas SET peserta_terdaftar = GREATEST(peserta_terdaftar - 1, 0),
			       status = CASE WHEN status='penuh' THEN 'aktif' ELSE status END WHERE id=$1`, kelasID); err != nil {
			return SettleResult{}, err
		}
		if err := tx.QueryRow(ctx, `SELECT id FROM pendaftaran WHERE user_id=$1 AND kelas_id=$2`, userID, kelasID).Scan(&pid); err != nil {
			return SettleResult{}, err
		}
	}

	if _, err := tx.Exec(ctx, `
		UPDATE transaksi SET status='sukses', midtrans_txn_id=$1, metode_pembayaran=NULLIF($2,''), pendaftaran_id=$3
		WHERE midtrans_order_id=$4`, txnID, metode, pid, orderID); err != nil {
		return SettleResult{}, err
	}
	if err := tx.Commit(ctx); err != nil {
		return SettleResult{}, err
	}
	return SettleResult{PendaftaranID: pid}, nil
}

func (r *transaksiRepo) MarkGagal(ctx context.Context, orderID string) error {
	_, err := r.db.Exec(ctx, `UPDATE transaksi SET status='gagal' WHERE midtrans_order_id=$1 AND status='pending'`, orderID)
	return err
}

func (r *transaksiRepo) MarkRefunded(ctx context.Context, orderID string) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var status string
	var pid *string
	err = tx.QueryRow(ctx, `SELECT status, pendaftaran_id FROM transaksi WHERE midtrans_order_id=$1 FOR UPDATE`, orderID).Scan(&status, &pid)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	if err != nil {
		return err
	}
	if status == "refund" {
		return nil
	}
	if pid != nil {
		var kelasID string
		if err := tx.QueryRow(ctx, `SELECT kelas_id FROM pendaftaran WHERE id=$1`, *pid).Scan(&kelasID); err == nil {
			if _, err := tx.Exec(ctx, `
				UPDATE kelas SET peserta_terdaftar = GREATEST(peserta_terdaftar - 1, 0),
				       status = CASE WHEN status='penuh' THEN 'aktif' ELSE status END WHERE id=$1`, kelasID); err != nil {
				return err
			}
			if _, err := tx.Exec(ctx, `UPDATE pendaftaran SET status='dibatalkan' WHERE id=$1`, *pid); err != nil {
				return err
			}
		}
	}
	if _, err := tx.Exec(ctx, `UPDATE transaksi SET status='refund' WHERE midtrans_order_id=$1`, orderID); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (r *transaksiRepo) UpdateStatusAdmin(ctx context.Context, id, status string) error {
	tag, err := r.db.Exec(ctx, `UPDATE transaksi SET status=$1 WHERE id=$2`, status, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
