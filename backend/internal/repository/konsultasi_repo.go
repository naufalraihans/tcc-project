package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"tcc-itpln/backend/internal/domain"
)

type konsultasiRepo struct{ db *pgxpool.Pool }

func NewKonsultasiRepository(db *pgxpool.Pool) KonsultasiRepository { return &konsultasiRepo{db} }

const konsultasiCols = `id, user_id, nama_pengirim, topik_konsultasi, pesan, kontak, status, COALESCE(balasan,''), admin_id, created_at, updated_at`

func scanKonsultasi(row pgx.Row) (domain.Konsultasi, error) {
	var k domain.Konsultasi
	err := row.Scan(&k.ID, &k.UserID, &k.NamaPengirim, &k.TopikKonsultasi, &k.Pesan, &k.Kontak, &k.Status, &k.Balasan, &k.AdminID, &k.CreatedAt, &k.UpdatedAt)
	return k, err
}

func (r *konsultasiRepo) Create(ctx context.Context, k domain.Konsultasi) (domain.Konsultasi, error) {
	return scanKonsultasi(r.db.QueryRow(ctx, `
		INSERT INTO konsultasi (user_id, nama_pengirim, topik_konsultasi, pesan, kontak)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING `+konsultasiCols,
		k.UserID, k.NamaPengirim, k.TopikKonsultasi, k.Pesan, k.Kontak))
}

func (r *konsultasiRepo) list(ctx context.Context, where string, args ...any) ([]domain.Konsultasi, error) {
	rows, err := r.db.Query(ctx, `SELECT `+konsultasiCols+` FROM konsultasi `+where+` ORDER BY created_at DESC`, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []domain.Konsultasi{}
	for rows.Next() {
		k, err := scanKonsultasi(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, k)
	}
	return out, rows.Err()
}

func (r *konsultasiRepo) ListByUser(ctx context.Context, userID string) ([]domain.Konsultasi, error) {
	return r.list(ctx, `WHERE user_id = $1`, userID)
}

func (r *konsultasiRepo) ListAll(ctx context.Context, status string) ([]domain.Konsultasi, error) {
	if status != "" && status != "semua" {
		return r.list(ctx, `WHERE status = $1`, status)
	}
	return r.list(ctx, ``)
}

func (r *konsultasiRepo) GetByID(ctx context.Context, id string) (domain.Konsultasi, error) {
	k, err := scanKonsultasi(r.db.QueryRow(ctx, `SELECT `+konsultasiCols+` FROM konsultasi WHERE id=$1`, id))
	if errors.Is(err, pgx.ErrNoRows) {
		return k, ErrNotFound
	}
	return k, err
}

func (r *konsultasiRepo) UpdateAdmin(ctx context.Context, id, adminID, status, balasan string) (domain.Konsultasi, error) {
	k, err := scanKonsultasi(r.db.QueryRow(ctx, `
		UPDATE konsultasi SET status=$1, balasan=NULLIF($2,''), admin_id=$3 WHERE id=$4
		RETURNING `+konsultasiCols, status, balasan, adminID, id))
	if errors.Is(err, pgx.ErrNoRows) {
		return k, ErrNotFound
	}
	return k, err
}
