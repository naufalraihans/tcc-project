package middleware

import (
	"github.com/gin-gonic/gin"

	"tcc-itpln/backend/pkg/utils"
)

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString(CtxRole) != "admin" {
			utils.Err(c, 403, "FORBIDDEN", "akses ditolak")
			c.Abort()
			return
		}
		c.Next()
	}
}
