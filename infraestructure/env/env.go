package env

import (
	"os"
)

type Environment struct {
	MONGO_URL          string
	MONGO_DB_NAME      string
	PATIENT_COLLECTION string
	DOCTOR_COLLECTION  string
	MYSQL_URL          string
	MYSQL_DB_NAME      string
	MYSQL_USER         string
	MYSQL_PASSWORD     string
	JWT_KEY            string
}

var env Environment

func init() {
	env = Environment{
		MONGO_URL:          os.Getenv("MONGO_URL"),
		MONGO_DB_NAME:      os.Getenv("MONGO_DB_NAME"),
		PATIENT_COLLECTION: os.Getenv("PATIENT_COLLECTION"),
		DOCTOR_COLLECTION:  os.Getenv("DOCTOR_COLLECTION"),
		MYSQL_URL:          os.Getenv("MYSQL_URL"),
		MYSQL_DB_NAME:      os.Getenv("MYSQL_DB_NAME"),
		MYSQL_USER:         os.Getenv("MYSQL_USER"),
		MYSQL_PASSWORD:     os.Getenv("MYSQL_PASSWORD"),
		JWT_KEY:            os.Getenv("JWT_KEY"),
	}
}

func GetEnv() *Environment {
	return &env
}
