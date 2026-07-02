package dto

type ProfileRequest struct {
	FullName  string `json:"full_name" binding:"required"`
	Phone     string `json:"phone"`
	AvatarURL string `json:"avatar_url"`
}
