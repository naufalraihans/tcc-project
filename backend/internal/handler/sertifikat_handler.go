package handler

import (
	"github.com/gin-gonic/gin"

	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/middleware"
	"tcc-itpln/backend/internal/usecase"
	"tcc-itpln/backend/pkg/utils"
)

type SertifikatHandler struct{ uc usecase.SertifikatUsecase }

func NewSertifikatHandler(uc usecase.SertifikatUsecase) *SertifikatHandler {
	return &SertifikatHandler{uc}
}

func (h *SertifikatHandler) Issue(c *gin.Context) {
	var req dto.SertifikatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c, err)
		return
	}
	data, err := h.uc.Issue(c.Request.Context(), req.PendaftaranID)
	if err != nil {
		fail(c, err)
		return
	}
	utils.Created(c, data)
}

func (h *SertifikatHandler) ListSaya(c *gin.Context) {
	data, err := h.uc.ListSaya(c.Request.Context(), c.GetString(middleware.CtxUserID))
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *SertifikatHandler) Verify(c *gin.Context) {
	data, err := h.uc.Verify(c.Request.Context(), c.Param("nomor"))
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}
