package domain

import "time"

type MiniTopik struct {
	ID   string `json:"id"`
	Nama string `json:"nama"`
	Slug string `json:"slug"`
}

type MiniInstruktur struct {
	ID      string `json:"id"`
	Nama    string `json:"nama"`
	FotoURL string `json:"foto_url"`
}

type KelasListItem struct {
	ID               string          `json:"id"`
	Judul            string          `json:"judul"`
	Slug             string          `json:"slug"`
	Topik            *MiniTopik      `json:"topik"`
	Instruktur       *MiniInstruktur `json:"instruktur"`
	Format           string          `json:"format"`
	TipeHarga        string          `json:"tipe_harga"`
	Harga            float64         `json:"harga"`
	JadwalMulai      *time.Time      `json:"jadwal_mulai"`
	JadwalSelesai    *time.Time      `json:"jadwal_selesai"`
	Kuota            int             `json:"kuota"`
	PesertaTerdaftar int             `json:"peserta_terdaftar"`
	Status           string          `json:"status"`
}

type Kelas struct {
	ID               string          `json:"id"`
	Judul            string          `json:"judul"`
	Slug             string          `json:"slug"`
	Deskripsi        string          `json:"deskripsi"`
	Silabus          string          `json:"silabus"`
	Topik            *MiniTopik      `json:"topik"`
	Instruktur       *MiniInstruktur `json:"instruktur"`
	Format           string          `json:"format"`
	TipeHarga        string          `json:"tipe_harga"`
	Harga            float64         `json:"harga"`
	JadwalMulai      *time.Time      `json:"jadwal_mulai"`
	JadwalSelesai    *time.Time      `json:"jadwal_selesai"`
	DurasiMenit      int             `json:"durasi_menit"`
	Kuota            int             `json:"kuota"`
	PesertaTerdaftar int             `json:"peserta_terdaftar"`
	Status           string          `json:"status"`
	Lokasi           string          `json:"lokasi"`
	LinkMeeting      string          `json:"link_meeting"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type PagedKelas struct {
	Items      []KelasListItem `json:"items"`
	Pagination Pagination      `json:"pagination"`
}
