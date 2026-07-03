package usecase

import (
	"context"

	"tcc-itpln/backend/internal/domain"
	"tcc-itpln/backend/internal/dto"
)

type TopikUsecase interface {
	List(ctx context.Context) ([]domain.Topik, error)
	GetBySlug(ctx context.Context, slug string) (domain.Topik, error)
	Create(ctx context.Context, req dto.TopikRequest) (domain.Topik, error)
	Update(ctx context.Context, id string, req dto.TopikRequest) (domain.Topik, error)
	Delete(ctx context.Context, id string) error
}

type InstrukturUsecase interface {
	List(ctx context.Context) ([]domain.Instruktur, error)
	GetByID(ctx context.Context, id string) (domain.Instruktur, error)
	Create(ctx context.Context, req dto.InstrukturRequest) (domain.Instruktur, error)
	Update(ctx context.Context, id string, req dto.InstrukturRequest) (domain.Instruktur, error)
	Delete(ctx context.Context, id string) error
}

type KelasUsecase interface {
	List(ctx context.Context, f dto.KelasFilter) (domain.PagedKelas, error)
	GetBySlug(ctx context.Context, slug string) (domain.Kelas, error)
	Create(ctx context.Context, req dto.KelasRequest) (domain.Kelas, error)
	Update(ctx context.Context, id string, req dto.KelasRequest) (domain.Kelas, error)
	UpdateStatus(ctx context.Context, id, status string) error
	Delete(ctx context.Context, id string) error
}

type AuthResult struct {
	Token    string `json:"token"`
	Role     string `json:"role"`
	FullName string `json:"full_name"`
}

type AuthUsecase interface {
	Register(ctx context.Context, req dto.RegisterRequest) (AuthResult, error)
	Login(ctx context.Context, req dto.LoginRequest) (AuthResult, error)
	Me(ctx context.Context, id string) (domain.Profile, error)
	UpdateProfile(ctx context.Context, id string, req dto.ProfileRequest) (domain.Profile, error)
}

type DaftarResult struct {
	Type          string `json:"type"`
	PendaftaranID string `json:"pendaftaran_id,omitempty"`
	TransaksiID   string `json:"transaksi_id,omitempty"`
	SnapToken     string `json:"snap_token,omitempty"`
	RedirectURL   string `json:"redirect_url,omitempty"`
	Message       string `json:"message,omitempty"`
}

type PendaftaranUsecase interface {
	Daftar(ctx context.Context, userID, email, kelasID string) (DaftarResult, error)
	ListSaya(ctx context.Context, userID, status string) ([]domain.PendaftaranItem, error)
	DetailSaya(ctx context.Context, userID, id string) (domain.Pendaftaran, error)
	ListAll(ctx context.Context, status string) ([]domain.PendaftaranItem, error)
	UpdateStatus(ctx context.Context, id, status string) error
}

type TransaksiUsecase interface {
	ListSaya(ctx context.Context, userID string) ([]domain.Transaksi, error)
	ListAll(ctx context.Context) ([]domain.Transaksi, error)
	HandleWebhook(ctx context.Context, p dto.MidtransWebhook) error
	UpdateStatus(ctx context.Context, id, status string) error
}

type SertifikatUsecase interface {
	Issue(ctx context.Context, pendaftaranID string) (domain.Sertifikat, error)
	ListSaya(ctx context.Context, userID string) ([]domain.Sertifikat, error)
	Verify(ctx context.Context, nomor string) (domain.SertifikatVerif, error)
}

type MateriUsecase interface {
	ListForUser(ctx context.Context, slug, userID string) ([]domain.Materi, error)
	Create(ctx context.Context, kelasID string, req dto.MateriRequest) (domain.Materi, error)
	Update(ctx context.Context, id string, req dto.MateriRequest) (domain.Materi, error)
	Delete(ctx context.Context, id string) error
}

type KonsultasiUsecase interface {
	Create(ctx context.Context, userID string, req dto.KonsultasiRequest) (domain.Konsultasi, error)
	ListSaya(ctx context.Context, userID string) ([]domain.Konsultasi, error)
	ListAll(ctx context.Context, status string) ([]domain.Konsultasi, error)
	Detail(ctx context.Context, id string) (domain.Konsultasi, error)
	Respond(ctx context.Context, id, adminID string, req dto.KonsultasiAdminRequest) (domain.Konsultasi, error)
}

type GamifikasiUsecase interface {
	Beranda(ctx context.Context, userID string) (dto.DashboardResponse, error)
	Progress(ctx context.Context, userID string) (dto.ProgressResponse, error)
	MisiHariIni(ctx context.Context, userID string) (dto.MisiRingkasan, error)
	ListMisi(ctx context.Context) ([]domain.Misi, error)
	CreateMisi(ctx context.Context, req dto.MisiRequest) (domain.Misi, error)
	UpdateMisi(ctx context.Context, id string, req dto.MisiRequest) (domain.Misi, error)
	DeleteMisi(ctx context.Context, id string) error
}

type PengumumanUsecase interface {
	ListActive(ctx context.Context, tipe string) ([]domain.Pengumuman, error)
	ListAll(ctx context.Context) ([]domain.Pengumuman, error)
	Create(ctx context.Context, req dto.PengumumanRequest) (domain.Pengumuman, error)
	Update(ctx context.Context, id string, req dto.PengumumanRequest) (domain.Pengumuman, error)
	Delete(ctx context.Context, id string) error
}
