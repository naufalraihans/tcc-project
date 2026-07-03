package handler

import (
	"github.com/gin-gonic/gin"

	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/middleware"
	"tcc-itpln/backend/internal/usecase"
	"tcc-itpln/backend/pkg/utils"
)

type GamifikasiHandler struct{ uc usecase.GamifikasiUsecase }

func NewGamifikasiHandler(uc usecase.GamifikasiUsecase) *GamifikasiHandler {
	return &GamifikasiHandler{uc}
}

func (h *GamifikasiHandler) Beranda(c *gin.Context) {
	data, err := h.uc.Beranda(c.Request.Context(), c.GetString(middleware.CtxUserID))
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *GamifikasiHandler) Progress(c *gin.Context) {
	data, err := h.uc.Progress(c.Request.Context(), c.GetString(middleware.CtxUserID))
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *GamifikasiHandler) Misi(c *gin.Context) {
	data, err := h.uc.MisiHariIni(c.Request.Context(), c.GetString(middleware.CtxUserID))
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

// ── Admin misi ──

func (h *GamifikasiHandler) ListMisi(c *gin.Context) {
	data, err := h.uc.ListMisi(c.Request.Context())
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *GamifikasiHandler) CreateMisi(c *gin.Context) {
	var req dto.MisiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c, err)
		return
	}
	data, err := h.uc.CreateMisi(c.Request.Context(), req)
	if err != nil {
		fail(c, err)
		return
	}
	utils.Created(c, data)
}

func (h *GamifikasiHandler) UpdateMisi(c *gin.Context) {
	var req dto.MisiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c, err)
		return
	}
	data, err := h.uc.UpdateMisi(c.Request.Context(), c.Param("id"), req)
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *GamifikasiHandler) DeleteMisi(c *gin.Context) {
	if err := h.uc.DeleteMisi(c.Request.Context(), c.Param("id")); err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, gin.H{"deleted": true})
}

// ── Pengumuman ──

type PengumumanHandler struct{ uc usecase.PengumumanUsecase }

func NewPengumumanHandler(uc usecase.PengumumanUsecase) *PengumumanHandler {
	return &PengumumanHandler{uc}
}

func (h *PengumumanHandler) ListActive(c *gin.Context) {
	data, err := h.uc.ListActive(c.Request.Context(), c.Query("tipe"))
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *PengumumanHandler) ListAll(c *gin.Context) {
	data, err := h.uc.ListAll(c.Request.Context())
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, data)
}

func (h *PengumumanHandler) Create(c *gin.Context) {
	var req dto.PengumumanRequest
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

func (h *PengumumanHandler) Update(c *gin.Context) {
	var req dto.PengumumanRequest
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

func (h *PengumumanHandler) Delete(c *gin.Context) {
	if err := h.uc.Delete(c.Request.Context(), c.Param("id")); err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, gin.H{"deleted": true})
}
