package handler

import (
	"github.com/gin-gonic/gin"

	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/usecase"
	"tcc-itpln/backend/pkg/utils"
)

type InstrukturHandler struct{ uc usecase.InstrukturUsecase }

func NewInstrukturHandler(uc usecase.InstrukturUsecase) *InstrukturHandler {
	return &InstrukturHandler{uc}
}

func (h *InstrukturHandler) List(c *gin.Context) {
	data, err := h.uc.List(c.Request.Context())
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *InstrukturHandler) Detail(c *gin.Context) {
	data, err := h.uc.GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *InstrukturHandler) Create(c *gin.Context) {
	var req dto.InstrukturRequest
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

func (h *InstrukturHandler) Update(c *gin.Context) {
	var req dto.InstrukturRequest
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

func (h *InstrukturHandler) Delete(c *gin.Context) {
	if err := h.uc.Delete(c.Request.Context(), c.Param("id")); err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, nil)
}
