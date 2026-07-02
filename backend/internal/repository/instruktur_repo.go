package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"tcc-itpln/backend/internal/domain"
)

type instrukturRepo struct{ db *pgxpool.Pool }

func NewInstrukturRepository(db *pgxpool.Pool) InstrukturRepository { return &instrukturRepo{db} }

func (r *instrukturRepo) List(ctx context.Context) ([]domain.Instruktur, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, nama, COALESCE(jabatan,''), COALESCE(foto_url,''), COALESCE(bio,''), created_at
		FROM instruktur ORDER BY nama`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []domain.Instruktur{}
	for rows.Next() {
		var in domain.Instruktur
		if err := rows.Scan(&in.ID, &in.Nama, &in.Jabatan, &in.FotoURL, &in.Bio, &in.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, in)
	}
	return out, rows.Err()
}

func (r *instrukturRepo) GetByID(ctx context.Context, id string) (domain.Instruktur, error) {
	var in domain.Instruktur
	err := r.db.QueryRow(ctx, `
		SELECT id, nama, COALESCE(jabatan,''), COALESCE(foto_url,''), COALESCE(bio,''), created_at
		FROM instruktur WHERE id=$1`, id).
		Scan(&in.ID, &in.Nama, &in.Jabatan, &in.FotoURL, &in.Bio, &in.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return in, ErrNotFound
	}
	return in, err
}

func (r *instrukturRepo) Create(ctx context.Context, in domain.Instruktur) (domain.Instruktur, error) {
	err := r.db.QueryRow(ctx, `
		INSERT INTO instruktur (nama, jabatan, foto_url, bio)
		VALUES ($1, NULLIF($2,''), NULLIF($3,''), NULLIF($4,''))
		RETURNING id, created_at`,
		in.Nama, in.Jabatan, in.FotoURL, in.Bio).Scan(&in.ID, &in.CreatedAt)
	return in, err
}

func (r *instrukturRepo) Update(ctx context.Context, id string, in domain.Instruktur) (domain.Instruktur, error) {
	err := r.db.QueryRow(ctx, `
		UPDATE instruktur SET nama=$1, jabatan=NULLIF($2,''), foto_url=NULLIF($3,''), bio=NULLIF($4,'')
		WHERE id=$5
		RETURNING id, nama, COALESCE(jabatan,''), COALESCE(foto_url,''), COALESCE(bio,''), created_at`,
		in.Nama, in.Jabatan, in.FotoURL, in.Bio, id).
		Scan(&in.ID, &in.Nama, &in.Jabatan, &in.FotoURL, &in.Bio, &in.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return in, ErrNotFound
	}
	return in, err
}

func (r *instrukturRepo) Delete(ctx context.Context, id string) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM instruktur WHERE id=$1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
