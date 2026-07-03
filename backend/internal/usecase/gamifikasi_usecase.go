package usecase

import (
	"context"

	"tcc-itpln/backend/internal/domain"
	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/repository"
)

type gamifikasiUsecase struct {
	gam     repository.GamifikasiRepository
	peng    repository.PengumumanRepository
	profile repository.ProfileRepository
}

func NewGamifikasiUsecase(gam repository.GamifikasiRepository, peng repository.PengumumanRepository, profile repository.ProfileRepository) GamifikasiUsecase {
	return &gamifikasiUsecase{gam, peng, profile}
}

func (u *gamifikasiUsecase) buildProgress(ctx context.Context, userID string) (dto.ProgressResponse, error) {
	p, err := u.gam.GetProgress(ctx, userID)
	if err != nil {
		return dto.ProgressResponse{}, err
	}
	week, err := u.gam.WeeklyActivity(ctx, userID)
	if err != nil {
		return dto.ProgressResponse{}, err
	}
	aktif := 0
	for _, d := range week {
		if d.Aktif {
			aktif++
		}
	}
	return dto.ProgressResponse{
		XP:                  p.XP,
		Level:               p.Level,
		XPKeLevelBerikutnya: 100 - (p.XP % 100),
		StreakSaatIni:       p.StreakSaatIni,
		StreakTerpanjang:    p.StreakTerpanjang,
		HariAktifMingguIni:  aktif,
		AktivitasMinggu:     week,
	}, nil
}

func (u *gamifikasiUsecase) misiRingkasan(ctx context.Context, userID string) (dto.MisiRingkasan, error) {
	items, err := u.gam.ListMisiHariIni(ctx, userID)
	if err != nil {
		return dto.MisiRingkasan{}, err
	}
	selesai := 0
	for _, m := range items {
		if m.Selesai {
			selesai++
		}
	}
	return dto.MisiRingkasan{Selesai: selesai, Total: len(items), Items: items}, nil
}

func (u *gamifikasiUsecase) Beranda(ctx context.Context, userID string) (dto.DashboardResponse, error) {
	// mencatat aktivitas hari ini (streak) sebagai bagian dari load beranda
	if err := u.gam.RecordActivity(ctx, userID); err != nil {
		return dto.DashboardResponse{}, err
	}

	prof, err := u.profile.GetByID(ctx, userID)
	if err != nil {
		return dto.DashboardResponse{}, err
	}
	progress, err := u.buildProgress(ctx, userID)
	if err != nil {
		return dto.DashboardResponse{}, err
	}
	misi, err := u.misiRingkasan(ctx, userID)
	if err != nil {
		return dto.DashboardResponse{}, err
	}
	peng, err := u.peng.ListActive(ctx, "")
	if err != nil {
		return dto.DashboardResponse{}, err
	}

	return dto.DashboardResponse{
		Profil:     dto.ProfilRingkas{FullName: prof.FullName, AvatarURL: prof.AvatarURL},
		Progress:   progress,
		Misi:       misi,
		Pengumuman: peng,
	}, nil
}

func (u *gamifikasiUsecase) Progress(ctx context.Context, userID string) (dto.ProgressResponse, error) {
	return u.buildProgress(ctx, userID)
}

func (u *gamifikasiUsecase) MisiHariIni(ctx context.Context, userID string) (dto.MisiRingkasan, error) {
	return u.misiRingkasan(ctx, userID)
}

// ── Admin ───────────────────────────────

func (u *gamifikasiUsecase) ListMisi(ctx context.Context) ([]domain.Misi, error) {
	return u.gam.ListMisi(ctx)
}

func aktifOrTrue(v *bool) bool { return v == nil || *v }

func (u *gamifikasiUsecase) CreateMisi(ctx context.Context, req dto.MisiRequest) (domain.Misi, error) {
	return u.gam.CreateMisi(ctx, domain.Misi{
		Kode: req.Kode, Judul: req.Judul, Deskripsi: req.Deskripsi, Tipe: req.Tipe,
		Target: req.Target, XPReward: req.XPReward, Aktif: aktifOrTrue(req.Aktif),
	})
}

func (u *gamifikasiUsecase) UpdateMisi(ctx context.Context, id string, req dto.MisiRequest) (domain.Misi, error) {
	return u.gam.UpdateMisi(ctx, id, domain.Misi{
		Kode: req.Kode, Judul: req.Judul, Deskripsi: req.Deskripsi, Tipe: req.Tipe,
		Target: req.Target, XPReward: req.XPReward, Aktif: aktifOrTrue(req.Aktif),
	})
}

func (u *gamifikasiUsecase) DeleteMisi(ctx context.Context, id string) error {
	return u.gam.DeleteMisi(ctx, id)
}

// ── Pengumuman ───────────────────────────────

type pengumumanUsecase struct {
	repo repository.PengumumanRepository
}

func NewPengumumanUsecase(repo repository.PengumumanRepository) PengumumanUsecase {
	return &pengumumanUsecase{repo}
}

func (u *pengumumanUsecase) ListActive(ctx context.Context, tipe string) ([]domain.Pengumuman, error) {
	return u.repo.ListActive(ctx, tipe)
}

func (u *pengumumanUsecase) ListAll(ctx context.Context) ([]domain.Pengumuman, error) {
	return u.repo.ListAll(ctx)
}

func (u *pengumumanUsecase) Create(ctx context.Context, req dto.PengumumanRequest) (domain.Pengumuman, error) {
	return u.repo.Create(ctx, pengumumanFromReq(req))
}

func (u *pengumumanUsecase) Update(ctx context.Context, id string, req dto.PengumumanRequest) (domain.Pengumuman, error) {
	return u.repo.Update(ctx, id, pengumumanFromReq(req))
}

func (u *pengumumanUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}

func pengumumanFromReq(req dto.PengumumanRequest) domain.Pengumuman {
	return domain.Pengumuman{
		Judul: req.Judul, Isi: req.Isi, Tipe: req.Tipe, LabelAksi: req.LabelAksi,
		URLAksi: req.URLAksi, Urutan: req.Urutan, Aktif: aktifOrTrue(req.Aktif),
		Mulai: req.Mulai, Selesai: req.Selesai,
	}
}
