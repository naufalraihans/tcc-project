package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"tcc-itpln/backend/internal/domain"
)

type pengumumanRepo struct{ db *pgxpool.Pool }

func NewPengumumanRepository(db *pgxpool.Pool) PengumumanRepository { return &pengumumanRepo{db} }

const pengumumanCols = `id, judul, COALESCE(isi,''), tipe, COALESCE(label_aksi,''), COALESCE(url_aksi,''), urutan, aktif, mulai, selesai, created_at`

func scanPengumuman(row pgx.Row) (domain.Pengumuman, error) {
	var p domain.Pengumuman
	err := row.Scan(&p.ID, &p.Judul, &p.Isi, &p.Tipe, &p.LabelAksi, &p.URLAksi, &p.Urutan, &p.Aktif, &p.Mulai, &p.Selesai, &p.CreatedAt)
	return p, err
}

func (r *pengumumanRepo) query(ctx context.Context, where string, args ...any) ([]domain.Pengumuman, error) {
	rows, err := r.db.Query(ctx, `SELECT `+pengumumanCols+` FROM pengumuman `+where, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []domain.Pengumuman{}
	for rows.Next() {
		p, err := scanPengumuman(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

func (r *pengumumanRepo) ListActive(ctx context.Context, tipe string) ([]domain.Pengumuman, error) {
	where := `WHERE aktif AND (mulai IS NULL OR mulai <= NOW()) AND (selesai IS NULL OR selesai >= NOW())`
	if tipe != "" {
		return r.query(ctx, where+` AND tipe = $1 ORDER BY urutan, created_at`, tipe)
	}
	return r.query(ctx, where+` ORDER BY urutan, created_at`)
}

func (r *pengumumanRepo) ListAll(ctx context.Context) ([]domain.Pengumuman, error) {
	return r.query(ctx, `ORDER BY urutan, created_at`)
}

func (r *pengumumanRepo) Create(ctx context.Context, p domain.Pengumuman) (domain.Pengumuman, error) {
	return scanPengumuman(r.db.QueryRow(ctx, `
		INSERT INTO pengumuman (judul, isi, tipe, label_aksi, url_aksi, urutan, aktif, mulai, selesai)
		VALUES ($1, NULLIF($2,''), $3, NULLIF($4,''), NULLIF($5,''), $6, $7, $8, $9)
		RETURNING `+pengumumanCols,
		p.Judul, p.Isi, p.Tipe, p.LabelAksi, p.URLAksi, p.Urutan, p.Aktif, p.Mulai, p.Selesai))
}

func (r *pengumumanRepo) Update(ctx context.Context, id string, p domain.Pengumuman) (domain.Pengumuman, error) {
	out, err := scanPengumuman(r.db.QueryRow(ctx, `
		UPDATE pengumuman SET judul=$1, isi=NULLIF($2,''), tipe=$3, label_aksi=NULLIF($4,''),
			url_aksi=NULLIF($5,''), urutan=$6, aktif=$7, mulai=$8, selesai=$9
		WHERE id=$10 RETURNING `+pengumumanCols,
		p.Judul, p.Isi, p.Tipe, p.LabelAksi, p.URLAksi, p.Urutan, p.Aktif, p.Mulai, p.Selesai, id))
	if errors.Is(err, pgx.ErrNoRows) {
		return out, ErrNotFound
	}
	return out, err
}

func (r *pengumumanRepo) Delete(ctx context.Context, id string) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM pengumuman WHERE id=$1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
