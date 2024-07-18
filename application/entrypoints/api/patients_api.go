package api

import (
	"encoding/json"
	"net/http"

	"github.com/chava.gnolasco/polaris/application/commons"
	"github.com/chava.gnolasco/polaris/application/domain/commands"
	apiModel "github.com/chava.gnolasco/polaris/application/entrypoints/model"
	"github.com/chava.gnolasco/polaris/infraestructure/log"
	"go.uber.org/zap"
)

func GetPatiens(writer http.ResponseWriter, request *http.Request) {

	log.Info("GetPatiens",
		zap.String("uuid", request.Header.Get("uuid")),
		zap.String("method", request.Method),
		zap.String("path", request.URL.Path),
	)

	patients := []apiModel.PatientDto{}
	patientManager := new(commands.PatientsManager)
	for _, patient := range patientManager.ConsultPatientsList() {
		patients = append(patients, *commons.GetPatientDtoFromModel(&patient))
	}

	PatientsResponse := commons.CreateAllPatientsResponse(patients)

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(PatientsResponse)
}
