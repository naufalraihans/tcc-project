package repository

import (
	"context"
	"time"

	"tcc-itpln/backend/internal/domain"
	"tcc-itpln/backend/internal/dto"
)

type errSentinel string

func (e errSentinel) Error() string { return string(e) }

const (
	ErrNotFound     = errSentinel("not found")
	ErrKuotaPenuh   = errSentinel("kuota penuh")
	ErrSudahDaftar  = errSentinel("sudah terdaftar")
	ErrNotEnrolled  = errSentinel("belum terdaftar")
)

type SettleResult struct {
	AlreadyDone   bool
	KuotaFull     bool
	PendaftaranID string
}

type TopikRepository interface {
	List(ctx context.Context) ([]domain.Topik, error)
	GetBySlug(ctx context.Context, slug string) (domain.Topik, error)
	Create(ctx context.Context, t domain.Topik) (domain.Topik, error)
	Update(ctx context.Context, id string, t domain.Topik) (domain.Topik, error)
	Delete(ctx context.Context, id string) error
}

type InstrukturRepository interface {
	List(ctx context.Context) ([]domain.Instruktur, error)
	GetByID(ctx context.Context, id string) (domain.Instruktur, error)
	Create(ctx context.Context, in domain.Instruktur) (domain.Instruktur, error)
	Update(ctx context.Context, id string, in domain.Instruktur) (domain.Instruktur, error)
	Delete(ctx context.Context, id string) error
}

type KelasRepository interface {
	List(ctx context.Context, f dto.KelasFilter) (domain.PagedKelas, error)
	GetBySlug(ctx context.Context, slug string) (domain.Kelas, error)
	Create(ctx context.Context, req dto.KelasRequest, slug string) (domain.Kelas, error)
	Update(ctx context.Context, id string, req dto.KelasRequest, slug string) (domain.Kelas, error)
	UpdateStatus(ctx context.Context, id, status string) error
	Delete(ctx context.Context, id string) error
}

type ProfileRepository interface {
	Role(ctx context.Context, id string) (string, error)
	GetByID(ctx context.Context, id string) (domain.Profile, error)
	Update(ctx context.Context, id, fullName, phone, avatarURL string) (domain.Profile, error)
}

type PendaftaranRepository interface {
	KelasInfo(ctx context.Context, kelasID string) (domain.KelasDaftarInfo, error)
	Exists(ctx context.Context, userID, kelasID string) (bool, error)
	ScheduleConflict(ctx context.Context, userID string, mulai, selesai time.Time) (string, bool, error)
	EnrollFree(ctx context.Context, userID, kelasID string) (string, error)
	ListByUser(ctx context.Context, userID, status string) ([]domain.PendaftaranItem, error)
	GetByUser(ctx context.Context, userID, id string) (domain.Pendaftaran, error)
	ListAll(ctx context.Context, status string) ([]domain.PendaftaranItem, error)
	UpdateStatus(ctx context.Context, id, status string) error
}

type KonsultasiRepository interface {
	Create(ctx context.Context, k domain.Konsultasi) (domain.Konsultasi, error)
	ListByUser(ctx context.Context, userID string) ([]domain.Konsultasi, error)
	ListAll(ctx context.Context, status string) ([]domain.Konsultasi, error)
	GetByID(ctx context.Context, id string) (domain.Konsultasi, error)
	UpdateAdmin(ctx context.Context, id, adminID, status, balasan string) (domain.Konsultasi, error)
}

type TransaksiRepository interface {
	Create(ctx context.Context, userID, kelasID, orderID string, jumlah float64) (string, error)
	GetByOrderID(ctx context.Context, orderID string) (domain.Transaksi, error)
	ListByUser(ctx context.Context, userID string) ([]domain.Transaksi, error)
	ListAll(ctx context.Context) ([]domain.Transaksi, error)
	Settle(ctx context.Context, orderID, txnID, metode string) (SettleResult, error)
	MarkGagal(ctx context.Context, orderID string) error
	MarkRefunded(ctx context.Context, orderID string) error
	UpdateStatusAdmin(ctx context.Context, id, status string) error
}

type SertifikatRepository interface {
	Issue(ctx context.Context, pendaftaranID string) (domain.Sertifikat, error)
	ListByUser(ctx context.Context, userID string) ([]domain.Sertifikat, error)
	GetByNomor(ctx context.Context, nomor string) (domain.SertifikatVerif, error)
}

type MateriRepository interface {
	ListForUser(ctx context.Context, slug, userID string) ([]domain.Materi, error)
	Create(ctx context.Context, kelasID string, m domain.Materi) (domain.Materi, error)
	Update(ctx context.Context, id string, m domain.Materi) (domain.Materi, error)
	Delete(ctx context.Context, id string) error
}
