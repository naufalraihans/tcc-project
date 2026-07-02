package domain

import "time"

type Transaksi struct {
	ID               string    `json:"id"`
	UserID           string    `json:"user_id"`
	KelasID          string    `json:"kelas_id"`
	PendaftaranID    *string   `json:"pendaftaran_id"`
	MidtransOrderID  string    `json:"midtrans_order_id"`
	MidtransTxnID    string    `json:"midtrans_txn_id"`
	Jumlah           float64   `json:"jumlah"`
	Status           string    `json:"status"`
	MetodePembayaran string    `json:"metode_pembayaran"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
