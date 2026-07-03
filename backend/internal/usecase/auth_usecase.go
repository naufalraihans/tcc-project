package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"tcc-itpln/backend/internal/domain"
	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/repository"
	"tcc-itpln/backend/pkg/supabase"
)

const tokenTTL = 30 * 24 * time.Hour

var ErrInvalidCredentials = errSentinelUC("email atau password salah")

type authUsecase struct {
	profileRepo repository.ProfileRepository
	authRepo    repository.AuthRepository
	gam         repository.GamifikasiRepository
	secret      string
}

func NewAuthUsecase(profileRepo repository.ProfileRepository, authRepo repository.AuthRepository, gam repository.GamifikasiRepository, secret string) AuthUsecase {
	return &authUsecase{profileRepo, authRepo, gam, secret}
}

func (u *authUsecase) Register(ctx context.Context, req dto.RegisterRequest) (AuthResult, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return AuthResult{}, err
	}
	id, role, err := u.authRepo.CreateUser(ctx, email, string(hash), req.FullName)
	if err != nil {
		return AuthResult{}, err
	}
	return u.issue(id, email, role, req.FullName)
}

func (u *authUsecase) Login(ctx context.Context, req dto.LoginRequest) (AuthResult, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))
	id, hash, role, err := u.authRepo.Credentials(ctx, email)
	if errors.Is(err, repository.ErrNotFound) {
		return AuthResult{}, ErrInvalidCredentials
	}
	if err != nil {
		return AuthResult{}, err
	}
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(req.Password)) != nil {
		return AuthResult{}, ErrInvalidCredentials
	}
	prof, _ := u.profileRepo.GetByID(ctx, id)
	return u.issue(id, email, role, prof.FullName)
}

func (u *authUsecase) issue(id, email, role, fullName string) (AuthResult, error) {
	token, err := supabase.Generate(u.secret, id, email, tokenTTL)
	if err != nil {
		return AuthResult{}, err
	}
	return AuthResult{Token: token, Role: role, FullName: fullName}, nil
}

func (u *authUsecase) Me(ctx context.Context, id string) (domain.Profile, error) {
	return u.profileRepo.GetByID(ctx, id)
}

func (u *authUsecase) UpdateProfile(ctx context.Context, id string, req dto.ProfileRequest) (domain.Profile, error) {
	prof, err := u.profileRepo.Update(ctx, id, req.FullName, req.Phone, req.AvatarURL)
	if err == nil && strings.TrimSpace(prof.FullName) != "" && strings.TrimSpace(prof.Phone) != "" {
		_ = u.gam.IncrementByKode(ctx, id, "lengkapi_profil") // ponytail: best-effort
	}
	return prof, err
}
