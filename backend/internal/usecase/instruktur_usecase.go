package usecase

import (
	"context"

	"tcc-itpln/backend/internal/domain"
	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/repository"
)

type instrukturUsecase struct{ repo repository.InstrukturRepository }

func NewInstrukturUsecase(repo repository.InstrukturRepository) InstrukturUsecase {
	return &instrukturUsecase{repo}
}

func (u *instrukturUsecase) List(ctx context.Context) ([]domain.Instruktur, error) {
	return u.repo.List(ctx)
}

func (u *instrukturUsecase) GetByID(ctx context.Context, id string) (domain.Instruktur, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *instrukturUsecase) Create(ctx context.Context, req dto.InstrukturRequest) (domain.Instruktur, error) {
	return u.repo.Create(ctx, domain.Instruktur{Nama: req.Nama, Jabatan: req.Jabatan, FotoURL: req.FotoURL, Bio: req.Bio})
}

func (u *instrukturUsecase) Update(ctx context.Context, id string, req dto.InstrukturRequest) (domain.Instruktur, error) {
	return u.repo.Update(ctx, id, domain.Instruktur{Nama: req.Nama, Jabatan: req.Jabatan, FotoURL: req.FotoURL, Bio: req.Bio})
}

func (u *instrukturUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
