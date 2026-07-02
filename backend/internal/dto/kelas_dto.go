package dto

import "time"

type KelasFilter struct {
	Topik  string
	Format string
	Harga  string
	Status string
	Page   int
	Limit  int
}

type KelasRequest struct {
	Judul         string     `json:"judul" binding:"required"`
	Slug          string     `json:"slug"`
	Deskripsi     string     `json:"deskripsi"`
	Silabus       string     `json:"silabus"`
	TopikID       *string    `json:"topik_id"`
	InstrukturID  *string    `json:"instruktur_id"`
	Format        string     `json:"format" binding:"required,oneof=online offline hybrid"`
	TipeHarga     string     `json:"tipe_harga" binding:"required,oneof=gratis berbayar"`
	Harga         float64    `json:"harga"`
	JadwalMulai   *time.Time `json:"jadwal_mulai"`
	JadwalSelesai *time.Time `json:"jadwal_selesai"`
	DurasiMenit   int        `json:"durasi_menit"`
	Kuota         int        `json:"kuota"`
	Lokasi        string     `json:"lokasi"`
	LinkMeeting   string     `json:"link_meeting"`
}

type StatusRequest struct {
	Status string `json:"status" binding:"required,oneof=aktif penuh selesai"`
}

type TopikRequest struct {
	Nama      string `json:"nama" binding:"required"`
	Slug      string `json:"slug"`
	Deskripsi string `json:"deskripsi"`
	IconURL   string `json:"icon_url"`
}

type InstrukturRequest struct {
	Nama    string `json:"nama" binding:"required"`
	Jabatan string `json:"jabatan"`
	FotoURL string `json:"foto_url"`
	Bio     string `json:"bio"`
}
