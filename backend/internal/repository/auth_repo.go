package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type authRepo struct{ db *pgxpool.Pool }

func NewAuthRepository(db *pgxpool.Pool) AuthRepository { return &authRepo{db} }

func (r *authRepo) CreateUser(ctx context.Context, email, passwordHash, fullName string) (string, string, error) {
	var id string
	err := r.db.QueryRow(ctx, `
		INSERT INTO auth.users (email, encrypted_password, raw_user_meta_data)
		VALUES ($1, $2, jsonb_build_object('full_name', $3::text))
		RETURNING id`, email, passwordHash, fullName).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return "", "", ErrEmailTaken
		}
		return "", "", err
	}
	return id, "user", nil
}

func (r *authRepo) Credentials(ctx context.Context, email string) (string, string, string, error) {
	var id, hash, role string
	err := r.db.QueryRow(ctx, `
		SELECT u.id, COALESCE(u.encrypted_password, ''), p.role
		FROM auth.users u JOIN public.profiles p ON p.id = u.id
		WHERE u.email = $1`, email).Scan(&id, &hash, &role)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", "", "", ErrNotFound
	}
	return id, hash, role, err
}
