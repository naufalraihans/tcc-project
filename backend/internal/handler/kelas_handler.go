package handler

import (
	"github.com/gin-gonic/gin"

	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/usecase"
	"tcc-itpln/backend/pkg/utils"
)

type KelasHandler struct{ uc usecase.KelasUsecase }

func NewKelasHandler(uc usecase.KelasUsecase) *KelasHandler { return &KelasHandler{uc} }

func (h *KelasHandler) List(c *gin.Context) {
	f := dto.KelasFilter{
		Topik:  c.Query("topik"),
		Format: c.Query("format"),
		Harga:  c.Query("harga"),
		Status: c.Query("status"),
		Page:   atoiDefault(c.Query("page"), 0),
		Limit:  atoiDefault(c.Query("limit"), 0),
	}
	data, err := h.uc.List(c.Request.Context(), f)
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *KelasHandler) Detail(c *gin.Context) {
	data, err := h.uc.GetBySlug(c.Request.Context(), c.Param("slug"))
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *KelasHandler) Create(c *gin.Context) {
	var req dto.KelasRequest
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

func (h *KelasHandler) Update(c *gin.Context) {
	var req dto.KelasRequest
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

func (h *KelasHandler) UpdateStatus(c *gin.Context) {
	var req dto.StatusRequest
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

func (h *KelasHandler) Delete(c *gin.Context) {
	if err := h.uc.Delete(c.Request.Context(), c.Param("id")); err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, nil)
}
