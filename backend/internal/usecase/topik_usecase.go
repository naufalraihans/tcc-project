package usecase

import (
	"context"

	"tcc-itpln/backend/internal/domain"
	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/repository"
	"tcc-itpln/backend/pkg/utils"
)

type topikUsecase struct{ repo repository.TopikRepository }

func NewTopikUsecase(repo repository.TopikRepository) TopikUsecase { return &topikUsecase{repo} }

func (u *topikUsecase) List(ctx context.Context) ([]domain.Topik, error) {
	return u.repo.List(ctx)
}

func (u *topikUsecase) GetBySlug(ctx context.Context, slug string) (domain.Topik, error) {
	return u.repo.GetBySlug(ctx, slug)
}

func (u *topikUsecase) Create(ctx context.Context, req dto.TopikRequest) (domain.Topik, error) {
	return u.repo.Create(ctx, domain.Topik{
		Nama:      req.Nama,
		Slug:      slugOr(req.Slug, req.Nama),
		Deskripsi: req.Deskripsi,
		IconURL:   req.IconURL,
	})
}

func (u *topikUsecase) Update(ctx context.Context, id string, req dto.TopikRequest) (domain.Topik, error) {
	return u.repo.Update(ctx, id, domain.Topik{
		Nama:      req.Nama,
		Slug:      slugOr(req.Slug, req.Nama),
		Deskripsi: req.Deskripsi,
		IconURL:   req.IconURL,
	})
}

func (u *topikUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}

func slugOr(slug, fallback string) string {
	if slug != "" {
		return utils.Slugify(slug)
	}
	return utils.Slugify(fallback)
}
