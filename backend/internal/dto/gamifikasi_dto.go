package dto

import (
	"time"

	"tcc-itpln/backend/internal/domain"
)

type ProgressResponse struct {
	XP                  int                `json:"xp"`
	Level               int                `json:"level"`
	XPKeLevelBerikutnya int                `json:"xp_ke_level_berikutnya"`
	StreakSaatIni       int                `json:"streak_saat_ini"`
	StreakTerpanjang    int                `json:"streak_terpanjang"`
	HariAktifMingguIni  int                `json:"hari_aktif_minggu_ini"`
	AktivitasMinggu     []domain.HariAktif `json:"aktivitas_minggu"`
}

type MisiRingkasan struct {
	Selesai int                  `json:"selesai"`
	Total   int                  `json:"total"`
	Items   []domain.MisiHariIni `json:"items"`
}

type ProfilRingkas struct {
	FullName  string `json:"full_name"`
	AvatarURL string `json:"avatar_url"`
}

type DashboardResponse struct {
	Profil     ProfilRingkas       `json:"profil"`
	Progress   ProgressResponse    `json:"progress"`
	Misi       MisiRingkasan       `json:"misi"`
	Pengumuman []domain.Pengumuman `json:"pengumuman"`
}

type MisiRequest struct {
	Kode      string `json:"kode" binding:"required"`
	Judul     string `json:"judul" binding:"required"`
	Deskripsi string `json:"deskripsi"`
	Tipe      string `json:"tipe" binding:"required,oneof=harian mingguan sekali"`
	Target    int    `json:"target" binding:"required,min=1"`
	XPReward  int    `json:"xp_reward" binding:"min=0"`
	Aktif     *bool  `json:"aktif"`
}

type PengumumanRequest struct {
	Judul     string     `json:"judul" binding:"required"`
	Isi       string     `json:"isi"`
	Tipe      string     `json:"tipe" binding:"required,oneof=banner info"`
	LabelAksi string     `json:"label_aksi"`
	URLAksi   string     `json:"url_aksi"`
	Urutan    int        `json:"urutan"`
	Aktif     *bool      `json:"aktif"`
	Mulai     *time.Time `json:"mulai"`
	Selesai   *time.Time `json:"selesai"`
}
