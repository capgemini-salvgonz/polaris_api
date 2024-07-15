package model

// Appointment struct
type Appointment struct {
	AppointmentID string  `json:"appointment_id"`
	Date          string  `json:"date"`
	Time          string  `json:"time"`
	Status        string  `json:"status"` //scheduled, confirmed, completed, canceled
	Reason        string  `json:"reason"`
	Notes         string  `json:"notes"`
	Doctor        Doctor  `json:"doctor_information"`
	Patient       Patient `json:"patient_information"`
}
