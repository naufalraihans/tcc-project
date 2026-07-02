package handler

import (
	"github.com/gin-gonic/gin"

	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/usecase"
	"tcc-itpln/backend/pkg/utils"
)

type TopikHandler struct{ uc usecase.TopikUsecase }

func NewTopikHandler(uc usecase.TopikUsecase) *TopikHandler { return &TopikHandler{uc} }

func (h *TopikHandler) List(c *gin.Context) {
	data, err := h.uc.List(c.Request.Context())
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *TopikHandler) Detail(c *gin.Context) {
	data, err := h.uc.GetBySlug(c.Request.Context(), c.Param("slug"))
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *TopikHandler) Create(c *gin.Context) {
	var req dto.TopikRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c, err)
		return
	}
	data, err := h.uc.Create(c.Request.Context(), req)
	if err != nil {
		fail(c, err)
		return
	}
	utils.Created(c, data)
}

func (h *TopikHandler) Update(c *gin.Context) {
	var req dto.TopikRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c, err)
		return
	}
	data, err := h.uc.Update(c.Request.Context(), c.Param("id"), req)
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *TopikHandler) Delete(c *gin.Context) {
	if err := h.uc.Delete(c.Request.Context(), c.Param("id")); err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, nil)
}
