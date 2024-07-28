// Description: This file contains the main function of the application to bootstrap.
package main

import (
	"net/http"

	"github.com/chava.gnolasco/polaris/application/entrypoints/api"
	"github.com/chava.gnolasco/polaris/infraestructure/log"

	"github.com/gorilla/mux"
)

var router *mux.Router

// init initializes the application.
func init() {
	log.Info("Application starting...")
	router = mux.NewRouter()
}

// main is the entry point of the application.
func main() {
	router.HandleFunc("/api/v1/polaris/patients", api.GetPatiens).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/polaris/patients/{id}", api.GetPatientById).Methods(http.MethodGet)

	router.HandleFunc("/api/v1/padmin/users", api.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/padmin/login", api.PadminUserLogin).Methods(http.MethodPost)

	error := http.ListenAndServe(":8080", router)
	if error != nil {
		log.Error("Error starting the application")
	}
}
