package dto

// UserRequestDto represents the user request data transfer object
type UserLoginRequestDto struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
