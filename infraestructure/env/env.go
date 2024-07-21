package env

import (
	"os"
)

type Environment struct {
	MONGO_URL          string
	MONGO_DB_NAME      string
	PATIENT_COLLECTION string
	DOCTOR_COLLECTION  string
}

var env Environment

func init() {
	env = Environment{
		MONGO_URL:          os.Getenv("MONGO_URL"),
		MONGO_DB_NAME:      os.Getenv("MONGO_DB_NAME"),
		PATIENT_COLLECTION: "patients",
		DOCTOR_COLLECTION:  "doctors",
	}
}

func GetEnv() *Environment {
	return &env
}
