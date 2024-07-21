package commons

import (
	"encoding/json"
	"net/http"

	"github.com/chava.gnolasco/polaris/application/entrypoints/model"
	"github.com/chava.gnolasco/polaris/infraestructure/log"
	"go.uber.org/zap"
)

/*
It logs the request
*/
func LogRequest(apiFunction string, request *http.Request) {
	log.Info(apiFunction,
		zap.String("uuid", request.Header.Get("uuid")),
		zap.String("method", request.Method),
		zap.String("path", request.URL.Path),
	)
}

/*
It creates an error response
*/
func ResponseError(writer http.ResponseWriter, code int, message string) {
	writer.WriteHeader(code)
	writer.Header().Set("Content-Type", "application/json")
	errorResponse := model.ErrorResponse{Code: code, Message: message}
	json.NewEncoder(writer).Encode(errorResponse)
}
