package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"tcc-itpln/backend/internal/domain"
)

type gamRepo struct{ db *pgxpool.Pool }

func NewGamifikasiRepository(db *pgxpool.Pool) GamifikasiRepository { return &gamRepo{db} }

// hariLabel: index 0 = Senin (sesuai date_trunc('week') Postgres)
var hariLabel = [7]string{"Sen", "Sel", "Rab", "Kam", "Jum", "Sab", "Min"}

// RecordActivity menandai user aktif hari ini; jika baru pertama hari ini,
// naikkan streak + XP (+5). Idempoten per hari via UNIQUE(user_id, tanggal).
func (r *gamRepo) RecordActivity(ctx context.Context, userID string) error {
	if _, err := r.db.Exec(ctx,
		`INSERT INTO user_progress (user_id) VALUES ($1) ON CONFLICT (user_id) DO NOTHING`, userID); err != nil {
		return err
	}
	tag, err := r.db.Exec(ctx,
		`INSERT INTO aktivitas_harian (user_id, tanggal, xp_didapat) VALUES ($1, CURRENT_DATE, 5)
		 ON CONFLICT (user_id, tanggal) DO NOTHING`, userID)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return nil // sudah aktif hari ini
	}
	_, err = r.db.Exec(ctx, `
		UPDATE user_progress SET
			streak_saat_ini   = CASE WHEN tanggal_aktif_terakhir = CURRENT_DATE - 1 THEN streak_saat_ini + 1 ELSE 1 END,
			streak_terpanjang = GREATEST(streak_terpanjang,
			                    CASE WHEN tanggal_aktif_terakhir = CURRENT_DATE - 1 THEN streak_saat_ini + 1 ELSE 1 END),
			xp    = xp + 5,
			level = 1 + (xp + 5) / 100,
			tanggal_aktif_terakhir = CURRENT_DATE,
			updated_at = NOW()
		WHERE user_id = $1`, userID)
	return err
}

func (r *gamRepo) GetProgress(ctx context.Context, userID string) (domain.UserProgress, error) {
	p := domain.UserProgress{UserID: userID, Level: 1}
	err := r.db.QueryRow(ctx,
		`SELECT xp, level, streak_saat_ini, streak_terpanjang, tanggal_aktif_terakhir
		 FROM user_progress WHERE user_id = $1`, userID).
		Scan(&p.XP, &p.Level, &p.StreakSaatIni, &p.StreakTerpanjang, &p.TanggalAktifTerakhir)
	if errors.Is(err, pgx.ErrNoRows) {
		return p, nil
	}
	return p, err
}

func (r *gamRepo) WeeklyActivity(ctx context.Context, userID string) ([]domain.HariAktif, error) {
	rows, err := r.db.Query(ctx, `
		SELECT d::date,
		       EXISTS (SELECT 1 FROM aktivitas_harian a WHERE a.user_id = $1 AND a.tanggal = d::date)
		FROM generate_series(date_trunc('week', CURRENT_DATE),
		                     date_trunc('week', CURRENT_DATE) + INTERVAL '6 day',
		                     INTERVAL '1 day') d
		ORDER BY d`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := make([]domain.HariAktif, 0, 7)
	i := 0
	for rows.Next() {
		var h domain.HariAktif
		var tgl time.Time
		var aktif bool
		if err := rows.Scan(&tgl, &aktif); err != nil {
			return nil, err
		}
		h.Tanggal = tgl.Format("2006-01-02")
		h.Aktif = aktif
		if i < 7 {
			h.Hari = hariLabel[i]
		}
		out = append(out, h)
		i++
	}
	return out, rows.Err()
}

func (r *gamRepo) ListMisiHariIni(ctx context.Context, userID string) ([]domain.MisiHariIni, error) {
	// lazy-create baris misi harian untuk hari ini
	if _, err := r.db.Exec(ctx, `
		INSERT INTO misi_user (user_id, misi_id, tanggal)
		SELECT $1, m.id, CURRENT_DATE FROM misi m WHERE m.aktif AND m.tipe = 'harian'
		ON CONFLICT (user_id, misi_id, tanggal) DO NOTHING`, userID); err != nil {
		return nil, err
	}
	rows, err := r.db.Query(ctx, `
		SELECT m.id, m.kode, m.judul, COALESCE(m.deskripsi,''), m.target, m.xp_reward, mu.progres, mu.selesai
		FROM misi_user mu JOIN misi m ON m.id = mu.misi_id
		WHERE mu.user_id = $1 AND mu.tanggal = CURRENT_DATE AND m.aktif AND m.tipe = 'harian'
		ORDER BY m.created_at`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []domain.MisiHariIni{}
	for rows.Next() {
		var m domain.MisiHariIni
		if err := rows.Scan(&m.ID, &m.Kode, &m.Judul, &m.Deskripsi, &m.Target, &m.XPReward, &m.Progres, &m.Selesai); err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	return out, rows.Err()
}

// IncrementByKode menaikkan progres misi (by kode) untuk user hari ini, dan
// memberi XP tepat sekali saat misi selesai. Idempoten & aman dipanggil berulang.
func (r *gamRepo) IncrementByKode(ctx context.Context, userID, kode string) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var misiID string
	var target, reward int
	err = tx.QueryRow(ctx,
		`SELECT id, target, xp_reward FROM misi WHERE kode = $1 AND aktif LIMIT 1`, kode).
		Scan(&misiID, &target, &reward)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil // misi tak ada / nonaktif → no-op
	}
	if err != nil {
		return err
	}

	var progres int
	var selesai bool
	err = tx.QueryRow(ctx, `
		INSERT INTO misi_user (user_id, misi_id, tanggal, progres, selesai, selesai_at)
		VALUES ($1, $2, CURRENT_DATE, 1, 1 >= $3, CASE WHEN 1 >= $3 THEN NOW() END)
		ON CONFLICT (user_id, misi_id, tanggal) DO UPDATE
			SET progres    = misi_user.progres + 1,
			    selesai    = (misi_user.progres + 1) >= $3,
			    selesai_at = CASE WHEN (misi_user.progres + 1) >= $3 THEN NOW() ELSE misi_user.selesai_at END
			WHERE misi_user.selesai = false
		RETURNING progres, selesai`, userID, misiID, target).Scan(&progres, &selesai)
	if errors.Is(err, pgx.ErrNoRows) {
		return tx.Commit(ctx) // sudah selesai sebelumnya → tidak ada perubahan
	}
	if err != nil {
		return err
	}

	if selesai { // baru selesai pada panggilan ini → beri XP sekali
		if _, err = tx.Exec(ctx,
			`INSERT INTO user_progress (user_id) VALUES ($1) ON CONFLICT (user_id) DO NOTHING`, userID); err != nil {
			return err
		}
		if _, err = tx.Exec(ctx,
			`UPDATE user_progress SET xp = xp + $2, level = 1 + (xp + $2) / 100, updated_at = NOW()
			 WHERE user_id = $1`, userID, reward); err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}

// AwardXP menambah XP langsung (mis. selesai kelas). Idempotency ada di pemanggil.
func (r *gamRepo) AwardXP(ctx context.Context, userID string, amount int) error {
	if _, err := r.db.Exec(ctx,
		`INSERT INTO user_progress (user_id) VALUES ($1) ON CONFLICT (user_id) DO NOTHING`, userID); err != nil {
		return err
	}
	_, err := r.db.Exec(ctx,
		`UPDATE user_progress SET xp = xp + $2, level = 1 + (xp + $2) / 100, updated_at = NOW()
		 WHERE user_id = $1`, userID, amount)
	return err
}

// ── Admin: kelola definisi misi ─────────────────────────────

const misiCols = `id, kode, judul, COALESCE(deskripsi,''), tipe, target, xp_reward, aktif, created_at`

func scanMisi(row pgx.Row) (domain.Misi, error) {
	var m domain.Misi
	err := row.Scan(&m.ID, &m.Kode, &m.Judul, &m.Deskripsi, &m.Tipe, &m.Target, &m.XPReward, &m.Aktif, &m.CreatedAt)
	return m, err
}

func (r *gamRepo) ListMisi(ctx context.Context) ([]domain.Misi, error) {
	rows, err := r.db.Query(ctx, `SELECT `+misiCols+` FROM misi ORDER BY created_at`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []domain.Misi{}
	for rows.Next() {
		m, err := scanMisi(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	return out, rows.Err()
}

func (r *gamRepo) CreateMisi(ctx context.Context, m domain.Misi) (domain.Misi, error) {
	return scanMisi(r.db.QueryRow(ctx, `
		INSERT INTO misi (kode, judul, deskripsi, tipe, target, xp_reward, aktif)
		VALUES ($1, $2, NULLIF($3,''), $4, $5, $6, $7)
		RETURNING `+misiCols,
		m.Kode, m.Judul, m.Deskripsi, m.Tipe, m.Target, m.XPReward, m.Aktif))
}

func (r *gamRepo) UpdateMisi(ctx context.Context, id string, m domain.Misi) (domain.Misi, error) {
	out, err := scanMisi(r.db.QueryRow(ctx, `
		UPDATE misi SET kode=$1, judul=$2, deskripsi=NULLIF($3,''), tipe=$4, target=$5, xp_reward=$6, aktif=$7
		WHERE id=$8 RETURNING `+misiCols,
		m.Kode, m.Judul, m.Deskripsi, m.Tipe, m.Target, m.XPReward, m.Aktif, id))
	if errors.Is(err, pgx.ErrNoRows) {
		return out, ErrNotFound
	}
	return out, err
}

func (r *gamRepo) DeleteMisi(ctx context.Context, id string) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM misi WHERE id=$1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
