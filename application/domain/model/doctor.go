package model

import (
	"time"
)

// Doctor struct represents the doctor model
type Doctor struct {
	DoctorId    string    `bson:"_id"` // echo "5543494708;ricardo.lopez@2code.com.mx" | base64
	Name        string    `bson:"name"`
	Specialty   string    `bson:"specialty"`
	Email       string    `bson:"email"`
	PhoneNumber string    `bson:"phone_number"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
	Clinic      Clinic    `bson:"clinic"`
}
