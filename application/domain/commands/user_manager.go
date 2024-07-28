package commands

import (
	"regexp"

	"github.com/chava.gnolasco/polaris/application/commons"
	"github.com/chava.gnolasco/polaris/application/domain/model"
	"github.com/chava.gnolasco/polaris/application/domain/ports"
)

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

func (um *UserManager) AuthenticatePadminUser(email string, password string) *model.User {

	user := um.ConsultUserByEmail(email)
	if user == nil || !commons.CheckPasswordHash(password, user.Password) {
		return nil
	}

	re := regexp.MustCompile(`\b(ADM|FS|AST)\b`)
	if !re.MatchString(user.Roles) {
		return nil
	}

	return user
}
