package middleware

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"

	"tcc-itpln/backend/pkg/supabase"
	"tcc-itpln/backend/pkg/utils"
)

const (
	CtxUserID = "user_id"
	CtxEmail  = "email"
	CtxRole   = "role"
)

func Auth(secret string, getRole func(context.Context, string) (string, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		token := strings.TrimPrefix(header, "Bearer ")
		if token == "" || token == header {
			utils.Err(c, 401, "UNAUTHORIZED", "token tidak ada")
			c.Abort()
			return
		}
		claims, err := supabase.ParseToken(secret, token)
		if err != nil {
			utils.Err(c, 401, "UNAUTHORIZED", "token tidak valid atau kedaluwarsa")
			c.Abort()
			return
		}
		role, err := getRole(c.Request.Context(), claims.Subject)
		if err != nil {
			utils.Err(c, 401, "UNAUTHORIZED", "profil tidak ditemukan")
			c.Abort()
			return
		}
		c.Set(CtxUserID, claims.Subject)
		c.Set(CtxEmail, claims.Email)
		c.Set(CtxRole, role)
		c.Next()
	}
}
