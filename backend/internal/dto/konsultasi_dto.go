package dto

type KonsultasiRequest struct {
	NamaPengirim    string `json:"nama_pengirim" binding:"required"`
	TopikKonsultasi string `json:"topik_konsultasi" binding:"required"`
	Pesan           string `json:"pesan" binding:"required"`
	Kontak          string `json:"kontak" binding:"required"`
}

type KonsultasiAdminRequest struct {
	Status  string `json:"status" binding:"required,oneof=menunggu diproses selesai ditolak"`
	Balasan string `json:"balasan"`
}
