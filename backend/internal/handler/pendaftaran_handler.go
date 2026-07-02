package handler

import (
	"errors"

	"github.com/gin-gonic/gin"

	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/middleware"
	"tcc-itpln/backend/internal/repository"
	"tcc-itpln/backend/internal/usecase"
	"tcc-itpln/backend/pkg/utils"
)

type PendaftaranHandler struct{ uc usecase.PendaftaranUsecase }

func NewPendaftaranHandler(uc usecase.PendaftaranUsecase) *PendaftaranHandler {
	return &PendaftaranHandler{uc}
}

func failDaftar(c *gin.Context, err error) {
	var ce usecase.ConflictError
	switch {
	case errors.Is(err, repository.ErrNotFound):
		utils.Err(c, 404, "NOT_FOUND", "kelas tidak ditemukan")
	case errors.Is(err, repository.ErrSudahDaftar):
		utils.Err(c, 409, "SUDAH_DAFTAR", "kamu sudah terdaftar di kelas ini")
	case errors.Is(err, repository.ErrKuotaPenuh):
		utils.Err(c, 409, "KUOTA_PENUH", "kuota kelas sudah penuh")
	case errors.Is(err, usecase.ErrKelasTidakAktif):
		utils.Err(c, 409, "KELAS_TIDAK_AKTIF", "kelas tidak menerima pendaftaran")
	case errors.As(err, &ce):
		utils.Err(c, 409, "SCHEDULE_CONFLICT", ce.Error())
	default:
		utils.Err(c, 500, "INTERNAL_ERROR", err.Error())
	}
}

func (h *PendaftaranHandler) Daftar(c *gin.Context) {
	var req dto.DaftarRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c, err)
		return
	}
	res, err := h.uc.Daftar(c.Request.Context(), c.GetString(middleware.CtxUserID), c.GetString(middleware.CtxEmail), req.KelasID)
	if err != nil {
		failDaftar(c, err)
		return
	}
	utils.OK(c, res)
}

func (h *PendaftaranHandler) ListSaya(c *gin.Context) {
	data, err := h.uc.ListSaya(c.Request.Context(), c.GetString(middleware.CtxUserID), c.Query("status"))
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *PendaftaranHandler) DetailSaya(c *gin.Context) {
	data, err := h.uc.DetailSaya(c.Request.Context(), c.GetString(middleware.CtxUserID), c.Param("id"))
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *PendaftaranHandler) ListAll(c *gin.Context) {
	data, err := h.uc.ListAll(c.Request.Context(), c.Query("status"))
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *PendaftaranHandler) UpdateStatus(c *gin.Context) {
	var req dto.PendaftaranStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c, err)
		return
	}
	if err := h.uc.UpdateStatus(c.Request.Context(), c.Param("id"), req.Status); err != nil {
		failDaftar(c, err)
		return
	}
	utils.OK(c, nil)
}
