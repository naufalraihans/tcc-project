package handler

import (
	"github.com/gin-gonic/gin"

	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/middleware"
	"tcc-itpln/backend/internal/usecase"
	"tcc-itpln/backend/pkg/utils"
)

type MateriHandler struct{ uc usecase.MateriUsecase }

func NewMateriHandler(uc usecase.MateriUsecase) *MateriHandler { return &MateriHandler{uc} }

func (h *MateriHandler) ListForUser(c *gin.Context) {
	data, err := h.uc.ListForUser(c.Request.Context(), c.Param("slug"), c.GetString(middleware.CtxUserID))
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *MateriHandler) Create(c *gin.Context) {
	var req dto.MateriRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c, err)
		return
	}
	data, err := h.uc.Create(c.Request.Context(), c.Param("id"), req)
	if err != nil {
		fail(c, err)
		return
	}
	utils.Created(c, data)
}

func (h *MateriHandler) Update(c *gin.Context) {
	var req dto.MateriRequest
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

func (h *MateriHandler) Delete(c *gin.Context) {
	if err := h.uc.Delete(c.Request.Context(), c.Param("id")); err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, nil)
}
