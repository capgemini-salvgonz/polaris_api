package commons

import (
	"github.com/chava.gnolasco/polaris/application/domain/model"
	"github.com/chava.gnolasco/polaris/application/entrypoints/dto"
)

func GetUserDtoFromModel(model *model.User, jwt string) *dto.UserDto {

	return &dto.UserDto{
		Id:          model.Id,
		Email:       model.Email,
		PhoneNumber: model.PhoneNumber,
		Roles:       model.Roles,
		CreatedAt:   model.CreatedAt,
		JWT:         jwt,
	}
}
