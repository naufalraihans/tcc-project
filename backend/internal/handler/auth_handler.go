package handler

import (
	"github.com/gin-gonic/gin"

	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/middleware"
	"tcc-itpln/backend/internal/usecase"
	"tcc-itpln/backend/pkg/utils"
)

type AuthHandler struct{ uc usecase.AuthUsecase }

func NewAuthHandler(uc usecase.AuthUsecase) *AuthHandler { return &AuthHandler{uc} }

func (h *AuthHandler) Me(c *gin.Context) {
	p, err := h.uc.Me(c.Request.Context(), c.GetString(middleware.CtxUserID))
	if err != nil {
		fail(c, err)
		return
	}
	p.Email = c.GetString(middleware.CtxEmail)
	utils.OK(c, p)
}

func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	var req dto.ProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c, err)
		return
	}
	p, err := h.uc.UpdateProfile(c.Request.Context(), c.GetString(middleware.CtxUserID), req)
	if err != nil {
		fail(c, err)
		return
	}
	p.Email = c.GetString(middleware.CtxEmail)
	utils.OK(c, p)
}
