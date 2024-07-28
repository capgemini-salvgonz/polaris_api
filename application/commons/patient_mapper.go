package commons

import (
	"github.com/chava.gnolasco/polaris/application/domain/model"
	"github.com/chava.gnolasco/polaris/application/entrypoints/dto"
)

func GetPatientDtoFromModel(model *model.Patient) *dto.PatientDto {

	return &dto.PatientDto{
		PatientId:   model.PatientId,
		Name:        model.Name,
		LastName:    model.LastName,
		DateOfBirth: model.DateOfBirth,
		Email:       model.Email,
		PhoneNumber: model.PhoneNumber,
		DoctorId:    model.DoctorId,
	}
}

func CreateAllPatientsResponse(patients []dto.PatientDto) *dto.PatientsResponse {
	PatientsResponse := dto.PatientsResponse{
		Timestamp: GetTimeStampt(),
		Patients:  patients,
	}

	return &PatientsResponse
}
