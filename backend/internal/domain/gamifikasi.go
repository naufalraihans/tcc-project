package domain

import "time"

type UserProgress struct {
	UserID               string
	XP                   int
	Level                int
	StreakSaatIni        int
	StreakTerpanjang     int
	TanggalAktifTerakhir *time.Time
}

type HariAktif struct {
	Tanggal string `json:"tanggal"` // YYYY-MM-DD
	Hari    string `json:"hari"`    // Sen..Min
	Aktif   bool   `json:"aktif"`
}

type Misi struct {
	ID        string    `json:"id"`
	Kode      string    `json:"kode"`
	Judul     string    `json:"judul"`
	Deskripsi string    `json:"deskripsi"`
	Tipe      string    `json:"tipe"`
	Target    int       `json:"target"`
	XPReward  int       `json:"xp_reward"`
	Aktif     bool      `json:"aktif"`
	CreatedAt time.Time `json:"created_at"`
}

// MisiHariIni = definisi misi + progres user hari ini (untuk "Misi Hari Ini").
type MisiHariIni struct {
	ID        string `json:"id"`
	Kode      string `json:"kode"`
	Judul     string `json:"judul"`
	Deskripsi string `json:"deskripsi"`
	Target    int    `json:"target"`
	XPReward  int    `json:"xp_reward"`
	Progres   int    `json:"progres"`
	Selesai   bool   `json:"selesai"`
}

type Pengumuman struct {
	ID        string     `json:"id"`
	Judul     string     `json:"judul"`
	Isi       string     `json:"isi"`
	Tipe      string     `json:"tipe"`
	LabelAksi string     `json:"label_aksi"`
	URLAksi   string     `json:"url_aksi"`
	Urutan    int        `json:"urutan"`
	Aktif     bool       `json:"aktif"`
	Mulai     *time.Time `json:"mulai"`
	Selesai   *time.Time `json:"selesai"`
	CreatedAt time.Time  `json:"created_at"`
}
