package usecase

import (
	"context"

	"tcc-itpln/backend/internal/domain"
	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/repository"
)

type konsultasiUsecase struct{ repo repository.KonsultasiRepository }

func NewKonsultasiUsecase(repo repository.KonsultasiRepository) KonsultasiUsecase {
	return &konsultasiUsecase{repo}
}

func (u *konsultasiUsecase) Create(ctx context.Context, userID string, req dto.KonsultasiRequest) (domain.Konsultasi, error) {
	return u.repo.Create(ctx, domain.Konsultasi{
		UserID:          userID,
		NamaPengirim:    req.NamaPengirim,
		TopikKonsultasi: req.TopikKonsultasi,
		Pesan:           req.Pesan,
		Kontak:          req.Kontak,
	})
}

func (u *konsultasiUsecase) ListSaya(ctx context.Context, userID string) ([]domain.Konsultasi, error) {
	return u.repo.ListByUser(ctx, userID)
}

func (u *konsultasiUsecase) ListAll(ctx context.Context, status string) ([]domain.Konsultasi, error) {
	return u.repo.ListAll(ctx, status)
}

func (u *konsultasiUsecase) Detail(ctx context.Context, id string) (domain.Konsultasi, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *konsultasiUsecase) Respond(ctx context.Context, id, adminID string, req dto.KonsultasiAdminRequest) (domain.Konsultasi, error) {
	return u.repo.UpdateAdmin(ctx, id, adminID, req.Status, req.Balasan)
}
