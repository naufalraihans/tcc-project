package domain

import "time"

type Instruktur struct {
	ID        string    `json:"id"`
	Nama      string    `json:"nama"`
	Jabatan   string    `json:"jabatan"`
	FotoURL   string    `json:"foto_url"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
}
