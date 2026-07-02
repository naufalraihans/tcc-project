package dto

type SertifikatRequest struct {
	PendaftaranID string `json:"pendaftaran_id" binding:"required"`
}
