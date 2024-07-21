package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/chava.gnolasco/polaris/application/adapters"
	"github.com/chava.gnolasco/polaris/application/commons"
	"github.com/chava.gnolasco/polaris/application/domain/commands"
	apiModel "github.com/chava.gnolasco/polaris/application/entrypoints/model"
	"github.com/chava.gnolasco/polaris/infraestructure/log"
)

/*
it gets the patient manager
*/
func getPatientManager(request *http.Request) *commands.PatientsManager {
	context := request.Context()
	patientManager := commands.PatientsManager{
		PatientRepository: adapters.NewPatientRepositoryAdapter(context),
	}
	return &patientManager
}

/*
it manages the request to get all patients for the API GET /api/v1/patients
It returns a list of patients
It returns an empty list if no patients are found
*/
func GetPatiens(writer http.ResponseWriter, request *http.Request) {

	commons.LogRequest("GetPatients", request)
	patientManager := getPatientManager(request)

	patients := []apiModel.PatientDto{}

	for _, patient := range patientManager.ConsultPatientsList() {
		patients = append(patients, *commons.GetPatientDtoFromModel(&patient))
	}

	PatientsResponse := commons.CreateAllPatientsResponse(patients)

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(PatientsResponse)
}

/*
it manages the request to get a patient by its id for the API GET /api/v1/patients/{id}
It returns a patient
It returns a 404 status code if no patient is found
*/
func GetPatientById(writer http.ResponseWriter, request *http.Request) {

	commons.LogRequest("GetPatientById", request)
	patientManager := getPatientManager(request)

	id := mux.Vars(request)["id"]
	log.Info("ID: " + id)
	patient := patientManager.ConsultPatientById(id)

	if patient == nil {
		commons.ResponseError(writer, http.StatusNotFound, "Patient not found")
		return
	}

	patientDto := commons.GetPatientDtoFromModel(patient)

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(patientDto)
}
