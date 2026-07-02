package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"tcc-itpln/backend/internal/domain"
)

type materiRepo struct{ db *pgxpool.Pool }

func NewMateriRepository(db *pgxpool.Pool) MateriRepository { return &materiRepo{db} }

func (r *materiRepo) ListForUser(ctx context.Context, slug, userID string) ([]domain.Materi, error) {
	var kelasID string
	err := r.db.QueryRow(ctx, `SELECT id FROM kelas WHERE slug=$1`, slug).Scan(&kelasID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	var one int
	err = r.db.QueryRow(ctx,
		`SELECT 1 FROM pendaftaran WHERE user_id=$1 AND kelas_id=$2 AND status IN ('aktif','selesai')`,
		userID, kelasID).Scan(&one)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotEnrolled
	}
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx,
		`SELECT id, kelas_id, judul, tipe, url, urutan, created_at FROM materi_kelas WHERE kelas_id=$1 ORDER BY urutan`, kelasID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []domain.Materi{}
	for rows.Next() {
		var m domain.Materi
		if err := rows.Scan(&m.ID, &m.KelasID, &m.Judul, &m.Tipe, &m.URL, &m.Urutan, &m.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	return out, rows.Err()
}

func (r *materiRepo) Create(ctx context.Context, kelasID string, m domain.Materi) (domain.Materi, error) {
	err := r.db.QueryRow(ctx, `
		INSERT INTO materi_kelas (kelas_id, judul, tipe, url, urutan) VALUES ($1,$2,$3,$4,$5)
		RETURNING id, kelas_id, judul, tipe, url, urutan, created_at`,
		kelasID, m.Judul, m.Tipe, m.URL, m.Urutan).
		Scan(&m.ID, &m.KelasID, &m.Judul, &m.Tipe, &m.URL, &m.Urutan, &m.CreatedAt)
	var pgErr interface{ SQLState() string }
	if errors.As(err, &pgErr) && pgErr.SQLState() == "23503" {
		return domain.Materi{}, ErrNotFound
	}
	return m, err
}

func (r *materiRepo) Update(ctx context.Context, id string, m domain.Materi) (domain.Materi, error) {
	err := r.db.QueryRow(ctx, `
		UPDATE materi_kelas SET judul=$1, tipe=$2, url=$3, urutan=$4 WHERE id=$5
		RETURNING id, kelas_id, judul, tipe, url, urutan, created_at`,
		m.Judul, m.Tipe, m.URL, m.Urutan, id).
		Scan(&m.ID, &m.KelasID, &m.Judul, &m.Tipe, &m.URL, &m.Urutan, &m.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.Materi{}, ErrNotFound
	}
	return m, err
}

func (r *materiRepo) Delete(ctx context.Context, id string) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM materi_kelas WHERE id=$1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
