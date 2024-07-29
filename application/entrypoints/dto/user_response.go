package dto

import (
	"time"
)

type UserDto struct {
	Id          int       `json:"user_id"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Roles       string    `json:"roles"`
	CreatedAt   time.Time `json:"created_at"`
	JWT         string    `json:"t_j"`
}

type UserResponse struct {
	Timestamp string    `json:"timestamp"`
	Users     []UserDto `json:"users"`
}
