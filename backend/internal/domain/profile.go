package domain

import "time"

type Profile struct {
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone"`
	AvatarURL string    `json:"avatar_url"`
	Role      string    `json:"role"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
