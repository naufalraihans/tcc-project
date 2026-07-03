package handler

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"tcc-itpln/backend/internal/dto"
	"tcc-itpln/backend/internal/usecase"
	"tcc-itpln/backend/pkg/utils"
)

type ChatHandler struct {
	uc   usecase.ChatUsecase
	mu   sync.Mutex
	hits []time.Time
}

func NewChatHandler(uc usecase.ChatUsecase) *ChatHandler {
	return &ChatHandler{uc: uc}
}

// ponytail: limiter global sederhana (maks 40/menit) untuk jaga kuota upstream Bynara (60/menit).
// Ganti ke per-IP kalau trafik naik.
func (h *ChatHandler) allow() bool {
	h.mu.Lock()
	defer h.mu.Unlock()
	cut := time.Now().Add(-time.Minute)
	kept := h.hits[:0]
	for _, t := range h.hits {
		if t.After(cut) {
			kept = append(kept, t)
		}
	}
	h.hits = kept
	if len(h.hits) >= 40 {
		return false
	}
	h.hits = append(h.hits, time.Now())
	return true
}

func (h *ChatHandler) Send(c *gin.Context) {
	if !h.allow() {
		utils.Err(c, 429, "RATE_LIMITED", "terlalu banyak permintaan, coba lagi beberapa saat")
		return
	}
	var req dto.ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c, err)
		return
	}
	reply, err := h.uc.Reply(c.Request.Context(), req.Messages)
	if err != nil {
		fail(c, err)
		return
	}
	utils.OK(c, dto.ChatResponse{Reply: reply})
}
