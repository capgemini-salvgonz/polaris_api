package adapters

import (
	"github.com/chava.gnolasco/polaris/application/domain/model"
	"github.com/chava.gnolasco/polaris/application/domain/ports"
)

type PatientRepositoryAdapter struct {
}

func NewPatientRepositoryAdapter() ports.PatientRepositoryPort {
	adapter := &PatientRepositoryAdapter{}
	return adapter
}

func (pra *PatientRepositoryAdapter) FindAll() []model.Patient {
	patients := []model.Patient{
		{
			PatientId:   "NTU0MzQ5NDcwOAo=",
			Name:        "Rick",
			LastName:    "Smith",
			DateOfBirth: "1980-05-29",
			Email:       "rick.smith@2code.com.mx",
			PhoneNumber: "5543590849",
		},
		{
			PatientId:   "NTU0MzU5MDg0OQo=",
			Name:        "John",
			LastName:    "Doe",
			DateOfBirth: "1982-11-11",
			Email:       "john.doe@2code.com.mx",
			PhoneNumber: "5543590849",
		},
	}

	return patients
}

func (pra *PatientRepositoryAdapter) FindById(id string) model.Patient {
	return model.Patient{}
}

func (pra *PatientRepositoryAdapter) FindByPhoneNumber(phoneNumber string) model.Patient {
	return model.Patient{}
}

func (pra *PatientRepositoryAdapter) FindByEmail(email string) model.Patient {
	return model.Patient{}
}

func (pra *PatientRepositoryAdapter) Save(patient model.Patient) model.Patient {
	return model.Patient{}
}
