package domain

import "time"

type Sertifikat struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	KelasID         string    `json:"kelas_id"`
	PendaftaranID   string    `json:"pendaftaran_id"`
	NomorSertifikat string    `json:"nomor_sertifikat"`
	URLSertifikat   string    `json:"url_sertifikat"`
	IssuedAt        time.Time `json:"issued_at"`
}

type SertifikatVerif struct {
	NomorSertifikat string    `json:"nomor_sertifikat"`
	NamaPenerima    string    `json:"nama_penerima"`
	Kelas           string    `json:"kelas"`
	IssuedAt        time.Time `json:"issued_at"`
	Valid           bool      `json:"valid"`
}
