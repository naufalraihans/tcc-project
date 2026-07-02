package usecase

import (
	"context"

	"tcc-itpln/backend/internal/domain"
	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/repository"
)

type materiUsecase struct{ repo repository.MateriRepository }

func NewMateriUsecase(repo repository.MateriRepository) MateriUsecase { return &materiUsecase{repo} }

func (u *materiUsecase) ListForUser(ctx context.Context, slug, userID string) ([]domain.Materi, error) {
	return u.repo.ListForUser(ctx, slug, userID)
}

func (u *materiUsecase) Create(ctx context.Context, kelasID string, req dto.MateriRequest) (domain.Materi, error) {
	return u.repo.Create(ctx, kelasID, domain.Materi{Judul: req.Judul, Tipe: req.Tipe, URL: req.URL, Urutan: req.Urutan})
}

func (u *materiUsecase) Update(ctx context.Context, id string, req dto.MateriRequest) (domain.Materi, error) {
	return u.repo.Update(ctx, id, domain.Materi{Judul: req.Judul, Tipe: req.Tipe, URL: req.URL, Urutan: req.Urutan})
}

func (u *materiUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
