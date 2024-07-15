package model

import (
	"time"
)

// Doctor struct represents the doctor model
type Doctor struct {
	DoctorId    string    `json:"doctor_id"`
	Name        string    `json:"name"`
	Specialty   string    `json:"specialty"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Clinic      Clinic    `json:"clinic"`
}
