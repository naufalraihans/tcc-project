package domain

import "time"

type MiniKelas struct {
	ID     string `json:"id"`
	Judul  string `json:"judul"`
	Slug   string `json:"slug"`
	Format string `json:"format"`
}

type MiniUser struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
}

type PendaftaranItem struct {
	PendaftaranID string    `json:"pendaftaran_id"`
	Kelas         MiniKelas `json:"kelas"`
	User          *MiniUser `json:"user,omitempty"`
	Status        string    `json:"status"`
	TanggalDaftar time.Time `json:"tanggal_daftar"`
}

type Pendaftaran struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	KelasID       string    `json:"kelas_id"`
	Status        string    `json:"status"`
	TanggalDaftar time.Time `json:"tanggal_daftar"`
}

type KelasDaftarInfo struct {
	Judul         string
	TipeHarga     string
	Status        string
	Harga         float64
	JadwalMulai   *time.Time
	JadwalSelesai *time.Time
}
