package commands

import (
	"github.com/chava.gnolasco/polaris/application/domain/model"
)

type PatientsManager struct {
}

func (pm *PatientsManager) ConsultPatientsList() []model.Patient {

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
