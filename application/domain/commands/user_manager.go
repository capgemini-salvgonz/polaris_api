package commands

import (
	"time"

	"github.com/chava.gnolasco/polaris/application/commons"
	"github.com/chava.gnolasco/polaris/application/domain/model"
	"github.com/chava.gnolasco/polaris/application/domain/ports"
	"github.com/chava.gnolasco/polaris/infraestructure/env"
	"github.com/chava.gnolasco/polaris/infraestructure/log"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(env.GetEnv().JWT_KEY)

type JWTClaims struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Roles       string `json:"roles"`
	jwt.StandardClaims
}

type UserManager struct {
	UserRepository ports.UserRepositoryPort
}

func (um *UserManager) ConsultUsersList() []model.User {
	users := um.UserRepository.FindAll()
	return users
}

func (um *UserManager) ConsultUserByEmail(email string) *model.User {
	user := um.UserRepository.FindByEmail(email)
	return user
}

func (um *UserManager) ConsultUserByPhoneNumber(phoneNumber string) *model.User {
	user := um.UserRepository.FindByPhoneNumber(phoneNumber)
	return user
}

/*
It creates a JWT for an authenticated user.
*/
func (um *UserManager) GenerateJWT(user *model.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &JWTClaims{
		Id:          user.Id,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Roles:       user.Roles,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

/*
AuthenticatePadminUser is a function that authenticates a user with the roles ADM, FS or AST.
*/
func (um *UserManager) AuthenticatePadminUser(email string, password string) (*model.User, string) {

	log.Info("AuthenticatePadminUser, [" + email + "]")

	user := um.ConsultUserByEmail(email)

	if user == nil || !commons.CheckPasswordHash(password, user.Password) {
		return nil, ""
	}

	log.Info("User found: " + user.PhoneNumber)
	jwt, _ := um.GenerateJWT(user)

	return user, jwt
}
