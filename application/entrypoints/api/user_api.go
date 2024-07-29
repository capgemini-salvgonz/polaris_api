package api

import (
	"encoding/json"
	"net/http"

	"github.com/chava.gnolasco/polaris/application/adapters"
	"github.com/chava.gnolasco/polaris/application/commons"
	"github.com/chava.gnolasco/polaris/application/domain/commands"
	"github.com/chava.gnolasco/polaris/application/entrypoints/dto"
)

var userManager commands.UserManager

func init() {
	userManager = commands.UserManager{
		UserRepository: adapters.NewUserRepositoryAdapter(),
	}
}

func GetUsers(writer http.ResponseWriter, request *http.Request) {
	commons.LogRequest("GetUsers", request)
}

/*
UserLogin is the endpoint for user login.
*/
func PadminUserLogin(writer http.ResponseWriter, request *http.Request) {

	commons.LogRequest("UserLogin", request)
	var userLoginRequest dto.UserLoginRequestDto

	if err := json.NewDecoder(request.Body).Decode(&userLoginRequest); err != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, jwt := userManager.AuthenticatePadminUser(
		commons.DecodeBase64(userLoginRequest.Email),
		commons.DecodeBase64(userLoginRequest.Password))

	if user == nil {
		commons.ResponseError(writer, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	response := commons.GetUserDtoFromModel(user, jwt)

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
