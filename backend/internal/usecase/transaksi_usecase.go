package usecase

import (
	"context"

	"tcc-itpln/backend/internal/domain"
	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/repository"
	"tcc-itpln/backend/pkg/midtrans"
)

var ErrInvalidSignature = errSentinelUC("signature tidak valid")

type transaksiUsecase struct {
	repo repository.TransaksiRepository
	gam  repository.GamifikasiRepository
	mid  *midtrans.Client
}

func NewTransaksiUsecase(repo repository.TransaksiRepository, gam repository.GamifikasiRepository, mid *midtrans.Client) TransaksiUsecase {
	return &transaksiUsecase{repo, gam, mid}
}

func (u *transaksiUsecase) ListSaya(ctx context.Context, userID string) ([]domain.Transaksi, error) {
	return u.repo.ListByUser(ctx, userID)
}

func (u *transaksiUsecase) ListAll(ctx context.Context) ([]domain.Transaksi, error) {
	return u.repo.ListAll(ctx)
}

func (u *transaksiUsecase) UpdateStatus(ctx context.Context, id, status string) error {
	return u.repo.UpdateStatusAdmin(ctx, id, status)
}

func (u *transaksiUsecase) HandleWebhook(ctx context.Context, p dto.MidtransWebhook) error {
	if !u.mid.VerifySignature(p.OrderID, p.StatusCode, p.GrossAmount, p.SignatureKey) {
		return ErrInvalidSignature
	}
	switch p.TransactionStatus {
	case "settlement", "capture":
		res, err := u.repo.Settle(ctx, p.OrderID, p.TransactionID, p.PaymentType)
		if err != nil {
			return err
		}
		if res.KuotaFull {
			_ = u.mid.Refund(p.OrderID)
			return u.repo.MarkRefunded(ctx, p.OrderID)
		}
		if !res.AlreadyDone { // enroll berbayar baru sukses → hook misi daftar_kelas
			if tx, err := u.repo.GetByOrderID(ctx, p.OrderID); err == nil {
				_ = u.gam.IncrementByKode(ctx, tx.UserID, "daftar_kelas") // ponytail: best-effort
			}
		}
		return nil
	case "deny", "cancel", "expire":
		return u.repo.MarkGagal(ctx, p.OrderID)
	case "refund", "partial_refund":
		return u.repo.MarkRefunded(ctx, p.OrderID)
	}
	return nil
}
