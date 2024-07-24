package models

type Permission struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}
