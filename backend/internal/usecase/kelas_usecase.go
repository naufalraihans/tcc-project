package usecase

import (
	"context"

	"tcc-itpln/backend/internal/domain"
	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/repository"
)

type kelasUsecase struct{ repo repository.KelasRepository }

func NewKelasUsecase(repo repository.KelasRepository) KelasUsecase { return &kelasUsecase{repo} }

func (u *kelasUsecase) List(ctx context.Context, f dto.KelasFilter) (domain.PagedKelas, error) {
	if f.Page < 1 {
		f.Page = 1
	}
	if f.Limit < 1 || f.Limit > 100 {
		f.Limit = 12
	}
	if f.Status == "" {
		f.Status = "aktif"
	}
	if f.Status == "semua" {
		f.Status = ""
	}
	return u.repo.List(ctx, f)
}

func (u *kelasUsecase) GetBySlug(ctx context.Context, slug string) (domain.Kelas, error) {
	return u.repo.GetBySlug(ctx, slug)
}

func (u *kelasUsecase) Create(ctx context.Context, req dto.KelasRequest) (domain.Kelas, error) {
	return u.repo.Create(ctx, req, slugOr(req.Slug, req.Judul))
}

func (u *kelasUsecase) Update(ctx context.Context, id string, req dto.KelasRequest) (domain.Kelas, error) {
	return u.repo.Update(ctx, id, req, slugOr(req.Slug, req.Judul))
}

func (u *kelasUsecase) UpdateStatus(ctx context.Context, id, status string) error {
	return u.repo.UpdateStatus(ctx, id, status)
}

func (u *kelasUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
