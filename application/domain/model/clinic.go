package model

// Clinic struct represents the clinic entity
type Clinic struct {
	Name        string `bson:"name"`
	Address     string `bson:"address"`
	PhoneNumber string `bson:"phone_number"`
	Email       string `bson:"email"`
}
