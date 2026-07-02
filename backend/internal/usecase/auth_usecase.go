package usecase

import (
	"context"

	"tcc-itpln/backend/internal/domain"
	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/repository"
)

type authUsecase struct{ repo repository.ProfileRepository }

func NewAuthUsecase(repo repository.ProfileRepository) AuthUsecase { return &authUsecase{repo} }

func (u *authUsecase) Me(ctx context.Context, id string) (domain.Profile, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *authUsecase) UpdateProfile(ctx context.Context, id string, req dto.ProfileRequest) (domain.Profile, error) {
	return u.repo.Update(ctx, id, req.FullName, req.Phone, req.AvatarURL)
}
