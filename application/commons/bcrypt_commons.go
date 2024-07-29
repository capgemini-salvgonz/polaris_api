package commons

import (
	"github.com/chava.gnolasco/polaris/infraestructure/log"
	"golang.org/x/crypto/bcrypt"
)

/*
HashPassword hashes the password.
*/
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

/*
CheckPasswordHash checks the password hash.
*/
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		log.Error("Error checking password hash: " + err.Error())
	}

	return err == nil
}
