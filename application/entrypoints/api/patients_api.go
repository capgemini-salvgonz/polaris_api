package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/chava.gnolasco/polaris/application/entrypoints/model"
	"github.com/chava.gnolasco/polaris/infraestructure/log"
	"go.uber.org/zap"
)

func GetPatiens(writer http.ResponseWriter, request *http.Request) {

	log.Info("GetPatiens",
		zap.String("uuid", request.Header.Get("uuid")),
		zap.String("method", request.Method),
		zap.String("path", request.URL.Path),
	)

	patients := []model.PatientDto{
		{
			PatientId:   "NTU0MzQ5NDcwOAo=",
			Name:        "Rick",
			LastName:    "Smith",
			DateOfBirth: "1980-05-29",
			Email:       "",
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

	PatientsResponse := model.PatientsResponse{
		Timestamp: time.Now().Format(time.RFC3339), // ISO 8601 format
		Patients:  patients,
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(PatientsResponse)
}
