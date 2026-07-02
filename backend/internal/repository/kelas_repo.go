package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"tcc-itpln/backend/internal/domain"
	"tcc-itpln/backend/internal/dto"
)

type kelasRepo struct{ db *pgxpool.Pool }

func NewKelasRepository(db *pgxpool.Pool) KelasRepository { return &kelasRepo{db} }

func (r *kelasRepo) List(ctx context.Context, f dto.KelasFilter) (domain.PagedKelas, error) {
	conds := []string{"1=1"}
	args := []any{}
	add := func(clause string, val any) {
		args = append(args, val)
		conds = append(conds, fmt.Sprintf(clause, len(args)))
	}
	if f.Status != "" {
		add("k.status=$%d", f.Status)
	}
	if f.Format != "" {
		add("k.format=$%d", f.Format)
	}
	if f.Harga != "" {
		add("k.tipe_harga=$%d", f.Harga)
	}
	if f.Topik != "" {
		add("t.slug=$%d", f.Topik)
	}
	where := strings.Join(conds, " AND ")

	out := domain.PagedKelas{Items: []domain.KelasListItem{}, Pagination: domain.Pagination{Page: f.Page, Limit: f.Limit}}
	if err := r.db.QueryRow(ctx,
		`SELECT COUNT(*) FROM kelas k LEFT JOIN topik t ON k.topik_id=t.id WHERE `+where, args...).
		Scan(&out.Pagination.Total); err != nil {
		return out, err
	}
	out.Pagination.TotalPages = (out.Pagination.Total + f.Limit - 1) / f.Limit

	q := fmt.Sprintf(`
		SELECT k.id, k.judul, k.slug,
		       t.id, t.nama, t.slug,
		       i.id, i.nama, COALESCE(i.foto_url,''),
		       k.format, k.tipe_harga, k.harga::float8, k.jadwal_mulai, k.jadwal_selesai,
		       k.kuota, k.peserta_terdaftar, k.status
		FROM kelas k
		LEFT JOIN topik t ON k.topik_id=t.id
		LEFT JOIN instruktur i ON k.instruktur_id=i.id
		WHERE %s ORDER BY k.created_at DESC LIMIT $%d OFFSET $%d`,
		where, len(args)+1, len(args)+2)
	args = append(args, f.Limit, (f.Page-1)*f.Limit)

	rows, err := r.db.Query(ctx, q, args...)
	if err != nil {
		return out, err
	}
	defer rows.Close()
	for rows.Next() {
		var it domain.KelasListItem
		var tID, tNama, tSlug, iID, iNama *string
		var iFoto string
		if err := rows.Scan(&it.ID, &it.Judul, &it.Slug,
			&tID, &tNama, &tSlug, &iID, &iNama, &iFoto,
			&it.Format, &it.TipeHarga, &it.Harga, &it.JadwalMulai, &it.JadwalSelesai,
			&it.Kuota, &it.PesertaTerdaftar, &it.Status); err != nil {
			return out, err
		}
		if tID != nil {
			it.Topik = &domain.MiniTopik{ID: *tID, Nama: *tNama, Slug: *tSlug}
		}
		if iID != nil {
			it.Instruktur = &domain.MiniInstruktur{ID: *iID, Nama: *iNama, FotoURL: iFoto}
		}
		out.Items = append(out.Items, it)
	}
	return out, rows.Err()
}

func (r *kelasRepo) fetch(ctx context.Context, col string, val any) (domain.Kelas, error) {
	var k domain.Kelas
	var tID, tNama, tSlug, iID, iNama *string
	var iFoto string
	err := r.db.QueryRow(ctx, `
		SELECT k.id, k.judul, k.slug, COALESCE(k.deskripsi,''), COALESCE(k.silabus,''),
		       t.id, t.nama, t.slug, i.id, i.nama, COALESCE(i.foto_url,''),
		       k.format, k.tipe_harga, k.harga::float8, k.jadwal_mulai, k.jadwal_selesai,
		       COALESCE(k.durasi_menit,0), k.kuota, k.peserta_terdaftar, k.status,
		       COALESCE(k.lokasi,''), COALESCE(k.link_meeting,''), k.created_at, k.updated_at
		FROM kelas k
		LEFT JOIN topik t ON k.topik_id=t.id
		LEFT JOIN instruktur i ON k.instruktur_id=i.id
		WHERE k.`+col+`=$1`, val).
		Scan(&k.ID, &k.Judul, &k.Slug, &k.Deskripsi, &k.Silabus,
			&tID, &tNama, &tSlug, &iID, &iNama, &iFoto,
			&k.Format, &k.TipeHarga, &k.Harga, &k.JadwalMulai, &k.JadwalSelesai,
			&k.DurasiMenit, &k.Kuota, &k.PesertaTerdaftar, &k.Status,
			&k.Lokasi, &k.LinkMeeting, &k.CreatedAt, &k.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return k, ErrNotFound
	}
	if err != nil {
		return k, err
	}
	if tID != nil {
		k.Topik = &domain.MiniTopik{ID: *tID, Nama: *tNama, Slug: *tSlug}
	}
	if iID != nil {
		k.Instruktur = &domain.MiniInstruktur{ID: *iID, Nama: *iNama, FotoURL: iFoto}
	}
	return k, nil
}

func (r *kelasRepo) GetBySlug(ctx context.Context, slug string) (domain.Kelas, error) {
	return r.fetch(ctx, "slug", slug)
}

func (r *kelasRepo) Create(ctx context.Context, req dto.KelasRequest, slug string) (domain.Kelas, error) {
	var id string
	err := r.db.QueryRow(ctx, `
		INSERT INTO kelas (judul, slug, deskripsi, silabus, topik_id, instruktur_id, format, tipe_harga, harga,
		                   jadwal_mulai, jadwal_selesai, durasi_menit, kuota, lokasi, link_meeting)
		VALUES ($1,$2,NULLIF($3,''),NULLIF($4,''),$5,$6,$7,$8,$9,$10,$11,$12,$13,NULLIF($14,''),NULLIF($15,''))
		RETURNING id`,
		req.Judul, slug, req.Deskripsi, req.Silabus, req.TopikID, req.InstrukturID, req.Format, req.TipeHarga, req.Harga,
		req.JadwalMulai, req.JadwalSelesai, req.DurasiMenit, req.Kuota, req.Lokasi, req.LinkMeeting).Scan(&id)
	if err != nil {
		return domain.Kelas{}, err
	}
	return r.fetch(ctx, "id", id)
}

func (r *kelasRepo) Update(ctx context.Context, id string, req dto.KelasRequest, slug string) (domain.Kelas, error) {
	tag, err := r.db.Exec(ctx, `
		UPDATE kelas SET judul=$1, slug=$2, deskripsi=NULLIF($3,''), silabus=NULLIF($4,''),
		       topik_id=$5, instruktur_id=$6, format=$7, tipe_harga=$8, harga=$9,
		       jadwal_mulai=$10, jadwal_selesai=$11, durasi_menit=$12, kuota=$13,
		       lokasi=NULLIF($14,''), link_meeting=NULLIF($15,'')
		WHERE id=$16`,
		req.Judul, slug, req.Deskripsi, req.Silabus, req.TopikID, req.InstrukturID, req.Format, req.TipeHarga, req.Harga,
		req.JadwalMulai, req.JadwalSelesai, req.DurasiMenit, req.Kuota, req.Lokasi, req.LinkMeeting, id)
	if err != nil {
		return domain.Kelas{}, err
	}
	if tag.RowsAffected() == 0 {
		return domain.Kelas{}, ErrNotFound
	}
	return r.fetch(ctx, "id", id)
}

func (r *kelasRepo) UpdateStatus(ctx context.Context, id, status string) error {
	tag, err := r.db.Exec(ctx, `UPDATE kelas SET status=$1 WHERE id=$2`, status, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *kelasRepo) Delete(ctx context.Context, id string) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM kelas WHERE id=$1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
