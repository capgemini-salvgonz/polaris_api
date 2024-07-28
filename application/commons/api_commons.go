package commons

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/chava.gnolasco/polaris/application/entrypoints/dto"
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
	errorResponse := dto.ErrorResponse{Code: code, Message: message}
	json.NewEncoder(writer).Encode(errorResponse)
}

/*
It creates a success response
*/
func GetTimeStampt() string {
	return time.Now().Format(time.RFC3339)
}

/*
It encodes a string to base64
*/
func DecodeBase64(data string) string {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Error("Error decoding the string", zap.Error(err))
		return ""
	}
	return strings.TrimSpace(string(decoded))
}
