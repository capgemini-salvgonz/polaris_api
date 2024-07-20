package commands

import (
	"github.com/chava.gnolasco/polaris/application/domain/model"
	"github.com/chava.gnolasco/polaris/application/domain/ports"
)

type PatientsManager struct {
	PatientRepository ports.PatientRepositoryPort
}

func (pm *PatientsManager) ConsultPatientsList() []model.Patient {

	patients := pm.PatientRepository.FindAll()
	return patients
}
