package dto

type MidtransWebhook struct {
	OrderID           string `json:"order_id"`
	TransactionID     string `json:"transaction_id"`
	TransactionStatus string `json:"transaction_status"`
	StatusCode        string `json:"status_code"`
	GrossAmount       string `json:"gross_amount"`
	PaymentType       string `json:"payment_type"`
	SignatureKey      string `json:"signature_key"`
}

type TransaksiStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=pending sukses gagal refund"`
}
