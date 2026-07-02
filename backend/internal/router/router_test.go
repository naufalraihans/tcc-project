package router

import (
	"testing"

	"github.com/gin-gonic/gin"

	"tcc-itpln/backend/config"
)

func TestNewRegistersWithoutConflict(t *testing.T) {
	gin.SetMode(gin.TestMode)
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("router.New panicked on route registration: %v", r)
		}
	}()
	if New(config.Config{}, nil) == nil {
		t.Fatal("expected engine, got nil")
	}
}
