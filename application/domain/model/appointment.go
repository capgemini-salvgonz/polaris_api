package model

// Appointment struct
type Appointment struct {
	AppointmentID   string `json:"appointment_id"`
	Date            string `json:"date"`
	Time            string `json:"time"`
	Status          string `json:"status"` //scheduled, confirmed, completed, canceled
	Reason          string `json:"reason"`
	Notes           string `json:"notes"`
	DoctorId        string `json:"doctor_id"`
	PatientId       string `json:"patient_id"`
	PatientFullName string `json:"patient_full_name"`
}
