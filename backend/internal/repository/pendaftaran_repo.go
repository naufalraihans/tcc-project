package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"tcc-itpln/backend/internal/domain"
)

type pendaftaranRepo struct{ db *pgxpool.Pool }

func NewPendaftaranRepository(db *pgxpool.Pool) PendaftaranRepository { return &pendaftaranRepo{db} }

func (r *pendaftaranRepo) KelasInfo(ctx context.Context, kelasID string) (domain.KelasDaftarInfo, error) {
	var info domain.KelasDaftarInfo
	err := r.db.QueryRow(ctx,
		`SELECT judul, tipe_harga, status, harga::float8, jadwal_mulai, jadwal_selesai FROM kelas WHERE id=$1`, kelasID).
		Scan(&info.Judul, &info.TipeHarga, &info.Status, &info.Harga, &info.JadwalMulai, &info.JadwalSelesai)
	if errors.Is(err, pgx.ErrNoRows) {
		return info, ErrNotFound
	}
	return info, err
}

func (r *pendaftaranRepo) Exists(ctx context.Context, userID, kelasID string) (bool, error) {
	var one int
	err := r.db.QueryRow(ctx, `SELECT 1 FROM pendaftaran WHERE user_id=$1 AND kelas_id=$2`, userID, kelasID).Scan(&one)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *pendaftaranRepo) ScheduleConflict(ctx context.Context, userID string, mulai, selesai time.Time) (string, bool, error) {
	var judul string
	err := r.db.QueryRow(ctx, `
		SELECT k.judul FROM pendaftaran p JOIN kelas k ON p.kelas_id=k.id
		WHERE p.user_id=$1 AND p.status='aktif' AND k.jadwal_mulai < $3 AND k.jadwal_selesai > $2
		LIMIT 1`, userID, mulai, selesai).Scan(&judul)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}
	return judul, true, nil
}

func (r *pendaftaranRepo) EnrollFree(ctx context.Context, userID, kelasID string) (string, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	tag, err := tx.Exec(ctx, `
		UPDATE kelas SET peserta_terdaftar = peserta_terdaftar + 1,
		       status = CASE WHEN kuota > 0 AND peserta_terdaftar + 1 >= kuota THEN 'penuh' ELSE status END
		WHERE id = $1 AND status = 'aktif' AND (kuota = 0 OR peserta_terdaftar < kuota)`, kelasID)
	if err != nil {
		return "", err
	}
	if tag.RowsAffected() == 0 {
		return "", ErrKuotaPenuh
	}

	var id string
	err = tx.QueryRow(ctx,
		`INSERT INTO pendaftaran (user_id, kelas_id, status) VALUES ($1, $2, 'aktif') RETURNING id`,
		userID, kelasID).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return "", ErrSudahDaftar
		}
		return "", err
	}
	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return id, nil
}

func (r *pendaftaranRepo) ListByUser(ctx context.Context, userID, status string) ([]domain.PendaftaranItem, error) {
	q := `SELECT p.id, k.id, k.judul, k.slug, k.format, p.status, p.tanggal_daftar
	      FROM pendaftaran p JOIN kelas k ON p.kelas_id = k.id WHERE p.user_id = $1`
	args := []any{userID}
	if status != "" && status != "semua" {
		q += ` AND p.status = $2`
		args = append(args, status)
	}
	q += ` ORDER BY p.tanggal_daftar DESC`

	rows, err := r.db.Query(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []domain.PendaftaranItem{}
	for rows.Next() {
		var it domain.PendaftaranItem
		if err := rows.Scan(&it.PendaftaranID, &it.Kelas.ID, &it.Kelas.Judul, &it.Kelas.Slug, &it.Kelas.Format, &it.Status, &it.TanggalDaftar); err != nil {
			return nil, err
		}
		out = append(out, it)
	}
	return out, rows.Err()
}

func (r *pendaftaranRepo) GetByUser(ctx context.Context, userID, id string) (domain.Pendaftaran, error) {
	var p domain.Pendaftaran
	err := r.db.QueryRow(ctx,
		`SELECT id, user_id, kelas_id, status, tanggal_daftar FROM pendaftaran WHERE id=$1 AND user_id=$2`, id, userID).
		Scan(&p.ID, &p.UserID, &p.KelasID, &p.Status, &p.TanggalDaftar)
	if errors.Is(err, pgx.ErrNoRows) {
		return p, ErrNotFound
	}
	return p, err
}

func (r *pendaftaranRepo) ListAll(ctx context.Context, status string) ([]domain.PendaftaranItem, error) {
	q := `SELECT p.id, k.id, k.judul, k.slug, k.format, p.status, p.tanggal_daftar, u.id, COALESCE(u.full_name,'')
	      FROM pendaftaran p JOIN kelas k ON p.kelas_id = k.id JOIN profiles u ON p.user_id = u.id WHERE 1=1`
	args := []any{}
	if status != "" && status != "semua" {
		args = append(args, status)
		q += ` AND p.status = $1`
	}
	q += ` ORDER BY p.tanggal_daftar DESC`

	rows, err := r.db.Query(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []domain.PendaftaranItem{}
	for rows.Next() {
		var it domain.PendaftaranItem
		var u domain.MiniUser
		if err := rows.Scan(&it.PendaftaranID, &it.Kelas.ID, &it.Kelas.Judul, &it.Kelas.Slug, &it.Kelas.Format, &it.Status, &it.TanggalDaftar, &u.ID, &u.FullName); err != nil {
			return nil, err
		}
		it.User = &u
		out = append(out, it)
	}
	return out, rows.Err()
}

func (r *pendaftaranRepo) UpdateStatus(ctx context.Context, id, newStatus string) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var cur, kelasID string
	err = tx.QueryRow(ctx, `SELECT status, kelas_id FROM pendaftaran WHERE id=$1 FOR UPDATE`, id).Scan(&cur, &kelasID)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound
	}
	if err != nil {
		return err
	}

	if cur == "aktif" && newStatus == "dibatalkan" {
		if _, err := tx.Exec(ctx, `
			UPDATE kelas SET peserta_terdaftar = GREATEST(peserta_terdaftar - 1, 0),
			       status = CASE WHEN status = 'penuh' THEN 'aktif' ELSE status END
			WHERE id = $1`, kelasID); err != nil {
			return err
		}
	}
	if cur != "aktif" && newStatus == "aktif" {
		tag, err := tx.Exec(ctx, `
			UPDATE kelas SET peserta_terdaftar = peserta_terdaftar + 1,
			       status = CASE WHEN kuota > 0 AND peserta_terdaftar + 1 >= kuota THEN 'penuh' ELSE status END
			WHERE id = $1 AND (kuota = 0 OR peserta_terdaftar < kuota)`, kelasID)
		if err != nil {
			return err
		}
		if tag.RowsAffected() == 0 {
			return ErrKuotaPenuh
		}
	}

	if _, err := tx.Exec(ctx, `
		UPDATE pendaftaran SET status = $1,
		       tanggal_selesai = CASE WHEN $1 = 'selesai' THEN NOW() ELSE tanggal_selesai END
		WHERE id = $2`, newStatus, id); err != nil {
		return err
	}
	return tx.Commit(ctx)
}
