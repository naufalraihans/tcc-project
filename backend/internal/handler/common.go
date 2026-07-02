package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"tcc-itpln/backend/internal/repository"
	"tcc-itpln/backend/pkg/utils"
)

func fail(c *gin.Context, err error) {
	switch {
	case errors.Is(err, repository.ErrNotFound):
		utils.Err(c, 404, "NOT_FOUND", "resource tidak ditemukan")
	case errors.Is(err, repository.ErrNotEnrolled):
		utils.Err(c, 403, "FORBIDDEN", "kamu belum terdaftar di kelas ini")
	default:
		utils.Err(c, 500, "INTERNAL_ERROR", err.Error())
	}
}

func badRequest(c *gin.Context, err error) {
	utils.Err(c, 422, "VALIDATION_ERROR", err.Error())
}

func atoiDefault(s string, def int) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	}
	return def
}
