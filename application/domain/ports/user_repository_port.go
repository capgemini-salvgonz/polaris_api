package ports

import (
	"github.com/chava.gnolasco/polaris/application/domain/model"
)

type UserRepositoryPort interface {
	FindAll() []model.User
	FindByEmail(email string) *model.User
	FindByPhoneNumber(phoneNumber string) *model.User
}
