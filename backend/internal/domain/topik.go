package domain

import "time"

type Topik struct {
	ID          string    `json:"id"`
	Nama        string    `json:"nama"`
	Slug        string    `json:"slug"`
	Deskripsi   string    `json:"deskripsi"`
	IconURL     string    `json:"icon_url"`
	JumlahKelas int       `json:"jumlah_kelas,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
