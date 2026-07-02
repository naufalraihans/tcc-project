package handler

import (
	"github.com/gin-gonic/gin"

	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/middleware"
	"tcc-itpln/backend/internal/usecase"
	"tcc-itpln/backend/pkg/utils"
)

type KonsultasiHandler struct{ uc usecase.KonsultasiUsecase }

func NewKonsultasiHandler(uc usecase.KonsultasiUsecase) *KonsultasiHandler {
	return &KonsultasiHandler{uc}
}

func (h *KonsultasiHandler) Create(c *gin.Context) {
	var req dto.KonsultasiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c, err)
		return
	}
	data, err := h.uc.Create(c.Request.Context(), c.GetString(middleware.CtxUserID), req)
	if err != nil {
		fail(c, err)
		return
	}
	utils.Created(c, data)
}

func (h *KonsultasiHandler) ListSaya(c *gin.Context) {
	data, err := h.uc.ListSaya(c.Request.Context(), c.GetString(middleware.CtxUserID))
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *KonsultasiHandler) ListAll(c *gin.Context) {
	data, err := h.uc.ListAll(c.Request.Context(), c.Query("status"))
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *KonsultasiHandler) Detail(c *gin.Context) {
	data, err := h.uc.Detail(c.Request.Context(), c.Param("id"))
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *KonsultasiHandler) Respond(c *gin.Context) {
	var req dto.KonsultasiAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c, err)
		return
	}
	data, err := h.uc.Respond(c.Request.Context(), c.Param("id"), c.GetString(middleware.CtxUserID), req)
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}
