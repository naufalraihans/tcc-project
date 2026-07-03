package usecase

import (
	"context"
	"time"

	"tcc-itpln/backend/internal/domain"
	"tcc-itpln/backend/internal/repository"
	"tcc-itpln/backend/pkg/midtrans"
	"tcc-itpln/backend/pkg/utils"
)

type ConflictError struct {
	Judul   string
	Mulai   *time.Time
	Selesai *time.Time
}

func (e ConflictError) Error() string { return "jadwal bentrok dengan kelas '" + e.Judul + "'" }

var ErrKelasTidakAktif = errSentinelUC("kelas tidak menerima pendaftaran")

type errSentinelUC string

func (e errSentinelUC) Error() string { return string(e) }

type pendaftaranUsecase struct {
	repo        repository.PendaftaranRepository
	txRepo      repository.TransaksiRepository
	profileRepo repository.ProfileRepository
	gam         repository.GamifikasiRepository
	mid         *midtrans.Client
}

func NewPendaftaranUsecase(repo repository.PendaftaranRepository, txRepo repository.TransaksiRepository, profileRepo repository.ProfileRepository, gam repository.GamifikasiRepository, mid *midtrans.Client) PendaftaranUsecase {
	return &pendaftaranUsecase{repo, txRepo, profileRepo, gam, mid}
}

func (u *pendaftaranUsecase) Daftar(ctx context.Context, userID, email, kelasID string) (DaftarResult, error) {
	info, err := u.repo.KelasInfo(ctx, kelasID)
	if err != nil {
		return DaftarResult{}, err
	}
	if info.Status != "aktif" {
		return DaftarResult{}, ErrKelasTidakAktif
	}

	exists, err := u.repo.Exists(ctx, userID, kelasID)
	if err != nil {
		return DaftarResult{}, err
	}
	if exists {
		return DaftarResult{}, repository.ErrSudahDaftar
	}

	if info.JadwalMulai != nil && info.JadwalSelesai != nil {
		judul, conflict, err := u.repo.ScheduleConflict(ctx, userID, *info.JadwalMulai, *info.JadwalSelesai)
		if err != nil {
			return DaftarResult{}, err
		}
		if conflict {
			return DaftarResult{}, ConflictError{Judul: judul, Mulai: info.JadwalMulai, Selesai: info.JadwalSelesai}
		}
	}

	if info.TipeHarga == "berbayar" {
		return u.startPayment(ctx, userID, email, kelasID, info)
	}

	id, err := u.repo.EnrollFree(ctx, userID, kelasID)
	if err != nil {
		return DaftarResult{}, err
	}
	_ = u.gam.IncrementByKode(ctx, userID, "daftar_kelas") // ponytail: best-effort
	return DaftarResult{Type: "gratis", PendaftaranID: id, Message: "Berhasil mendaftar kelas"}, nil
}

func (u *pendaftaranUsecase) startPayment(ctx context.Context, userID, email, kelasID string, info domain.KelasDaftarInfo) (DaftarResult, error) {
	orderID := utils.OrderID()
	txID, err := u.txRepo.Create(ctx, userID, kelasID, orderID, info.Harga)
	if err != nil {
		return DaftarResult{}, err
	}
	prof, _ := u.profileRepo.GetByID(ctx, userID)
	token, redirect, err := u.mid.CreateSnapToken(midtrans.SnapRequest{
		OrderID:     orderID,
		GrossAmount: info.Harga,
		Customer:    midtrans.Customer{FirstName: prof.FullName, Email: email},
		Item:        midtrans.Item{ID: kelasID, Price: info.Harga, Quantity: 1, Name: info.Judul},
	})
	if err != nil {
		return DaftarResult{}, err
	}
	return DaftarResult{Type: "berbayar", TransaksiID: txID, SnapToken: token, RedirectURL: redirect}, nil
}

func (u *pendaftaranUsecase) ListSaya(ctx context.Context, userID, status string) ([]domain.PendaftaranItem, error) {
	return u.repo.ListByUser(ctx, userID, status)
}

func (u *pendaftaranUsecase) DetailSaya(ctx context.Context, userID, id string) (domain.Pendaftaran, error) {
	return u.repo.GetByUser(ctx, userID, id)
}

func (u *pendaftaranUsecase) ListAll(ctx context.Context, status string) ([]domain.PendaftaranItem, error) {
	return u.repo.ListAll(ctx, status)
}

func (u *pendaftaranUsecase) UpdateStatus(ctx context.Context, id, status string) error {
	completedUserID, err := u.repo.UpdateStatus(ctx, id, status)
	if err != nil {
		return err
	}
	if completedUserID != "" {
		_ = u.gam.AwardXP(ctx, completedUserID, 100) // ponytail: best-effort, selesai kelas
	}
	return nil
}
