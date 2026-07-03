package dto

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Messages []ChatMessage `json:"messages" binding:"required,min=1"`
}

type ChatResponse struct {
	Reply string `json:"reply"`
}
