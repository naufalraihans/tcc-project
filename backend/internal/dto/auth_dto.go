package dto

type ProfileRequest struct {
	FullName  string `json:"full_name" binding:"required"`
	Phone     string `json:"phone"`
	AvatarURL string `json:"avatar_url"`
}

type RegisterRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
