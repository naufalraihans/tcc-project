package domain

import "time"

type Konsultasi struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	NamaPengirim    string    `json:"nama_pengirim"`
	TopikKonsultasi string    `json:"topik_konsultasi"`
	Pesan           string    `json:"pesan"`
	Kontak          string    `json:"kontak"`
	Status          string    `json:"status"`
	Balasan         string    `json:"balasan"`
	AdminID         *string   `json:"admin_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
