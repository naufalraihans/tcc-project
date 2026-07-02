package handler

import (
	"errors"

	"github.com/gin-gonic/gin"

	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/middleware"
	"tcc-itpln/backend/internal/usecase"
	"tcc-itpln/backend/pkg/utils"
)

type TransaksiHandler struct{ uc usecase.TransaksiUsecase }

func NewTransaksiHandler(uc usecase.TransaksiUsecase) *TransaksiHandler { return &TransaksiHandler{uc} }

func (h *TransaksiHandler) ListSaya(c *gin.Context) {
	data, err := h.uc.ListSaya(c.Request.Context(), c.GetString(middleware.CtxUserID))
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *TransaksiHandler) ListAll(c *gin.Context) {
	data, err := h.uc.ListAll(c.Request.Context())
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *TransaksiHandler) UpdateStatus(c *gin.Context) {
	var req dto.TransaksiStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c, err)
		return
	}
	if err := h.uc.UpdateStatus(c.Request.Context(), c.Param("id"), req.Status); err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, nil)
}

func (h *TransaksiHandler) Webhook(c *gin.Context) {
	var p dto.MidtransWebhook
	if err := c.ShouldBindJSON(&p); err != nil {
		badRequest(c, err)
		return
	}
	if err := h.uc.HandleWebhook(c.Request.Context(), p); err != nil {
		if errors.Is(err, usecase.ErrInvalidSignature) {
			utils.Err(c, 400, "BAD_SIGNATURE", "signature tidak valid")
			return
		}
		fail(c, err)
		return
	}
	utils.OK(c, gin.H{"status": "ok"})
}
