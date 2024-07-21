package model

// PatientDto is a struct that represents the patient data transfer object
type PatientDto struct {
	PatientId   string `json:"patient_id"`
	Name        string `json:"name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	DoctorId    string `json:"doctor_id"`
}

// PatientsResponse is a struct that represents the list of patients to response
type PatientsResponse struct {
	Timestamp string       `json:"timestamp"`
	Patients  []PatientDto `json:"patients"`
}
