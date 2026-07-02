package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"tcc-itpln/backend/internal/domain"
)

type sertifikatRepo struct{ db *pgxpool.Pool }

func NewSertifikatRepository(db *pgxpool.Pool) SertifikatRepository { return &sertifikatRepo{db} }

func (r *sertifikatRepo) Issue(ctx context.Context, pendaftaranID string) (domain.Sertifikat, error) {
	var s domain.Sertifikat
	var userID, kelasID, status string
	err := r.db.QueryRow(ctx, `SELECT user_id, kelas_id, status FROM pendaftaran WHERE id=$1`, pendaftaranID).
		Scan(&userID, &kelasID, &status)
	if errors.Is(err, pgx.ErrNoRows) {
		return s, ErrNotFound
	}
	if err != nil {
		return s, err
	}

	year := time.Now().Year()
	var count int
	if err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM sertifikat WHERE nomor_sertifikat LIKE $1`,
		fmt.Sprintf("TCC-%d-%%", year)).Scan(&count); err != nil {
		return s, err
	}
	nomor := fmt.Sprintf("TCC-%d-%04d", year, count+1)

	err = r.db.QueryRow(ctx, `
		INSERT INTO sertifikat (user_id, kelas_id, pendaftaran_id, nomor_sertifikat)
		VALUES ($1,$2,$3,$4)
		RETURNING id, user_id, kelas_id, pendaftaran_id, nomor_sertifikat, COALESCE(url_sertifikat,''), issued_at`,
		userID, kelasID, pendaftaranID, nomor).
		Scan(&s.ID, &s.UserID, &s.KelasID, &s.PendaftaranID, &s.NomorSertifikat, &s.URLSertifikat, &s.IssuedAt)
	return s, err
}

func (r *sertifikatRepo) ListByUser(ctx context.Context, userID string) ([]domain.Sertifikat, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, user_id, kelas_id, pendaftaran_id, nomor_sertifikat, COALESCE(url_sertifikat,''), issued_at
		FROM sertifikat WHERE user_id=$1 ORDER BY issued_at DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []domain.Sertifikat{}
	for rows.Next() {
		var s domain.Sertifikat
		if err := rows.Scan(&s.ID, &s.UserID, &s.KelasID, &s.PendaftaranID, &s.NomorSertifikat, &s.URLSertifikat, &s.IssuedAt); err != nil {
			return nil, err
		}
		out = append(out, s)
	}
	return out, rows.Err()
}

func (r *sertifikatRepo) GetByNomor(ctx context.Context, nomor string) (domain.SertifikatVerif, error) {
	var v domain.SertifikatVerif
	err := r.db.QueryRow(ctx, `
		SELECT s.nomor_sertifikat, COALESCE(u.full_name,''), k.judul, s.issued_at
		FROM sertifikat s JOIN profiles u ON s.user_id=u.id JOIN kelas k ON s.kelas_id=k.id
		WHERE s.nomor_sertifikat=$1`, nomor).
		Scan(&v.NomorSertifikat, &v.NamaPenerima, &v.Kelas, &v.IssuedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return v, ErrNotFound
	}
	if err != nil {
		return v, err
	}
	v.Valid = true
	return v, nil
}
