package model

// Patient struct
type Patient struct {
	PatientId   string `json:"patient_id"`
	Name        string `json:"name"`
	LastNme     string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
