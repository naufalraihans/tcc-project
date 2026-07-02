package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"tcc-itpln/backend/internal/domain"
)

type topikRepo struct{ db *pgxpool.Pool }

func NewTopikRepository(db *pgxpool.Pool) TopikRepository { return &topikRepo{db} }

func (r *topikRepo) List(ctx context.Context) ([]domain.Topik, error) {
	rows, err := r.db.Query(ctx, `
		SELECT t.id, t.nama, t.slug, COALESCE(t.deskripsi,''), COALESCE(t.icon_url,''),
		       (SELECT COUNT(*) FROM kelas k WHERE k.topik_id = t.id), t.created_at
		FROM topik t ORDER BY t.nama`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []domain.Topik{}
	for rows.Next() {
		var t domain.Topik
		if err := rows.Scan(&t.ID, &t.Nama, &t.Slug, &t.Deskripsi, &t.IconURL, &t.JumlahKelas, &t.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

func (r *topikRepo) GetBySlug(ctx context.Context, slug string) (domain.Topik, error) {
	var t domain.Topik
	err := r.db.QueryRow(ctx, `
		SELECT id, nama, slug, COALESCE(deskripsi,''), COALESCE(icon_url,''),
		       (SELECT COUNT(*) FROM kelas k WHERE k.topik_id = topik.id), created_at
		FROM topik WHERE slug = $1`, slug).
		Scan(&t.ID, &t.Nama, &t.Slug, &t.Deskripsi, &t.IconURL, &t.JumlahKelas, &t.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return t, ErrNotFound
	}
	return t, err
}

func (r *topikRepo) Create(ctx context.Context, t domain.Topik) (domain.Topik, error) {
	err := r.db.QueryRow(ctx, `
		INSERT INTO topik (nama, slug, deskripsi, icon_url)
		VALUES ($1, $2, NULLIF($3,''), NULLIF($4,''))
		RETURNING id, created_at`,
		t.Nama, t.Slug, t.Deskripsi, t.IconURL).Scan(&t.ID, &t.CreatedAt)
	return t, err
}

func (r *topikRepo) Update(ctx context.Context, id string, t domain.Topik) (domain.Topik, error) {
	err := r.db.QueryRow(ctx, `
		UPDATE topik SET nama=$1, slug=$2, deskripsi=NULLIF($3,''), icon_url=NULLIF($4,'')
		WHERE id=$5
		RETURNING id, nama, slug, COALESCE(deskripsi,''), COALESCE(icon_url,''), created_at`,
		t.Nama, t.Slug, t.Deskripsi, t.IconURL, id).
		Scan(&t.ID, &t.Nama, &t.Slug, &t.Deskripsi, &t.IconURL, &t.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return t, ErrNotFound
	}
	return t, err
}

func (r *topikRepo) Delete(ctx context.Context, id string) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM topik WHERE id=$1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
