package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"tcc-itpln/backend/internal/domain"
)

type profileRepo struct{ db *pgxpool.Pool }

func NewProfileRepository(db *pgxpool.Pool) ProfileRepository { return &profileRepo{db} }

func (r *profileRepo) Role(ctx context.Context, id string) (string, error) {
	var role string
	err := r.db.QueryRow(ctx, `SELECT role FROM profiles WHERE id=$1`, id).Scan(&role)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", ErrNotFound
	}
	return role, err
}

func (r *profileRepo) GetByID(ctx context.Context, id string) (domain.Profile, error) {
	var p domain.Profile
	err := r.db.QueryRow(ctx, `
		SELECT id, COALESCE(full_name,''), COALESCE(phone,''), COALESCE(avatar_url,''), role, created_at
		FROM profiles WHERE id=$1`, id).
		Scan(&p.ID, &p.FullName, &p.Phone, &p.AvatarURL, &p.Role, &p.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return p, ErrNotFound
	}
	return p, err
}

func (r *profileRepo) Update(ctx context.Context, id, fullName, phone, avatarURL string) (domain.Profile, error) {
	var p domain.Profile
	err := r.db.QueryRow(ctx, `
		UPDATE profiles SET full_name=$1, phone=NULLIF($2,''), avatar_url=NULLIF($3,'')
		WHERE id=$4
		RETURNING id, COALESCE(full_name,''), COALESCE(phone,''), COALESCE(avatar_url,''), role, created_at`,
		fullName, phone, avatarURL, id).
		Scan(&p.ID, &p.FullName, &p.Phone, &p.AvatarURL, &p.Role, &p.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return p, ErrNotFound
	}
	return p, err
}
