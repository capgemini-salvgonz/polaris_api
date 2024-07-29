// Description: This file contains the main function of the application to bootstrap.
package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/chava.gnolasco/polaris/application/entrypoints/api"
	"github.com/chava.gnolasco/polaris/application/entrypoints/middleware"
	"github.com/chava.gnolasco/polaris/infraestructure/log"

	"github.com/gorilla/mux"
)

var router *mux.Router

// init initializes the application.
func init() {
	log.Info("Application starting...")
	router = mux.NewRouter()
}

/*
Main function that starts the application.
*/
func main() {
	// Start the pprof server
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	// Public resources
	router.HandleFunc("/api/v1/padmin/login", api.PadminUserLogin).Methods(http.MethodPost)

	// Private resources

	// --- Patients
	router.Handle("/api/v1/polaris/patients", middleware.JWTAuth(http.HandlerFunc(api.GetPatiens))).Methods(http.MethodGet)
	router.Handle("/api/v1/polaris/patients/{id}", middleware.JWTAuth(http.HandlerFunc(api.GetPatientById))).Methods(http.MethodGet)

	// --- Users
	router.Handle("/api/v1/padmin/users", middleware.JWTAuth(http.HandlerFunc(api.GetUsers))).Methods(http.MethodGet)

	error := http.ListenAndServe(":8080", router)
	if error != nil {
		log.Error("Error starting the application")
	}
}
