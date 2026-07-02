package dto

type MateriRequest struct {
	Judul  string `json:"judul" binding:"required"`
	Tipe   string `json:"tipe" binding:"required,oneof=file link video"`
	URL    string `json:"url" binding:"required"`
	Urutan int    `json:"urutan"`
}
