package model

// Patient struct
type Patient struct {
	PatientId   string `bson:"_id"`
	Name        string `bson:"name"`
	LastName    string `bson:"last_name"`
	DateOfBirth string `bson:"date_of_birth"`
	Email       string `bson:"email"`
	PhoneNumber string `bson:"phone_number"`
	DoctorId    string `bson:"doctor_id"`
}
