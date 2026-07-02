package usecase

import (
	"context"

	"tcc-itpln/backend/internal/domain"
	"tcc-itpln/backend/internal/repository"
)

type sertifikatUsecase struct{ repo repository.SertifikatRepository }

func NewSertifikatUsecase(repo repository.SertifikatRepository) SertifikatUsecase {
	return &sertifikatUsecase{repo}
}

func (u *sertifikatUsecase) Issue(ctx context.Context, pendaftaranID string) (domain.Sertifikat, error) {
	return u.repo.Issue(ctx, pendaftaranID)
}

func (u *sertifikatUsecase) ListSaya(ctx context.Context, userID string) ([]domain.Sertifikat, error) {
	return u.repo.ListByUser(ctx, userID)
}

func (u *sertifikatUsecase) Verify(ctx context.Context, nomor string) (domain.SertifikatVerif, error) {
	return u.repo.GetByNomor(ctx, nomor)
}
