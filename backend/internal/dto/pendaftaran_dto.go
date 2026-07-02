package dto

type DaftarRequest struct {
	KelasID string `json:"kelas_id" binding:"required"`
}

type PendaftaranStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=aktif selesai dibatalkan"`
}
