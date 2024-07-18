package commons

import (
	"time"

	domain "github.com/chava.gnolasco/polaris/application/domain/model"
	apiModel "github.com/chava.gnolasco/polaris/application/entrypoints/model"
)

func GetPatientDtoFromModel(model *domain.Patient) (dto *apiModel.PatientDto) {

	dto = &apiModel.PatientDto{
		PatientId:   model.PatientId,
		Name:        model.Name,
		LastName:    model.LastName,
		DateOfBirth: model.DateOfBirth,
		Email:       model.Email,
		PhoneNumber: model.PhoneNumber,
	}

	return
}

func CreateAllPatientsResponse(patients []apiModel.PatientDto) *apiModel.PatientsResponse {
	PatientsResponse := apiModel.PatientsResponse{
		Timestamp: time.Now().Format(time.RFC3339), // ISO 8601 format
		Patients:  patients,
	}

	return &PatientsResponse
}
