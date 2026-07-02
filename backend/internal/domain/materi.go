package domain

import "time"

type Materi struct {
	ID        string    `json:"id"`
	KelasID   string    `json:"kelas_id"`
	Judul     string    `json:"judul"`
	Tipe      string    `json:"tipe"`
	URL       string    `json:"url"`
	Urutan    int       `json:"urutan"`
	CreatedAt time.Time `json:"created_at"`
}
