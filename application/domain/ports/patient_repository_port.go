package ports

import (
	"github.com/chava.gnolasco/polaris/application/domain/model"
)

type PatientRepositoryPort interface {
	FindAll() []model.Patient
	FindById(id string) *model.Patient
	FindByPhoneNumber(phoneNumber string) *model.Patient
	FindByEmail(email string) *model.Patient
	Save(patient *model.Patient) *model.Patient
}
