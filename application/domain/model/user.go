package model

import (
	"time"
)

// User struct
// User representa la tabla users en Go
type User struct {
	Id          int       `gorm:"column:id;primaryKey;autoIncrement"`
	Email       string    `gorm:"column:email;unique"`
	PhoneNumber string    `gorm:"column:phone_number"`
	Password    string    `gorm:"column:password"`
	Roles       string    `gorm:"column:roles"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
}
