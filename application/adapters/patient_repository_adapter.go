package adapters

import (
	"context"

	"github.com/chava.gnolasco/polaris/application/domain/model"
	"github.com/chava.gnolasco/polaris/application/domain/ports"
	"github.com/chava.gnolasco/polaris/infraestructure/env"
	"github.com/chava.gnolasco/polaris/infraestructure/log"
	"github.com/chava.gnolasco/polaris/infraestructure/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

type PatientRepositoryAdapter struct {
	context context.Context
}

func NewPatientRepositoryAdapter(ctx context.Context) ports.PatientRepositoryPort {
	adapter := PatientRepositoryAdapter{context: ctx}
	return &adapter
}

/*
It retrieve all patients from the database
It returns a slice of patients
It returns an empty slice if no patients are found
*/
func (adapter *PatientRepositoryAdapter) FindAll() (patients []model.Patient) {

	patients = []model.Patient{}
	patientCollection := mongodb.GetCollection(env.GetEnv().PATIENT_COLLECTION)
	cursor, err := patientCollection.Find(adapter.context, bson.M{})

	if err != nil {
		log.Error("Error finding patients: ", zap.String("ERROR", err.Error()))
		return
	}

	defer cursor.Close(adapter.context)

	for cursor.Next(adapter.context) {
		var patient model.Patient
		err := cursor.Decode(&patient)
		if err != nil {
			log.Error("Error reading patients: ", zap.String("ERROR", err.Error()))
			return
		}

		patients = append(patients, patient)
	}

	return
}

/*
It retrieves a patient by its id
It returns a pointer to a patient
It returns nil if no patient is found
*/
func (adapter *PatientRepositoryAdapter) FindById(id string) *model.Patient {
	patientCollection := mongodb.GetCollection(env.GetEnv().PATIENT_COLLECTION)

	var patient model.Patient
	err := patientCollection.FindOne(adapter.context, bson.M{"_id": id}).Decode(&patient)

	if err != nil {
		log.Error("Error finding patient by id: ", zap.String("ERROR", err.Error()))
		return nil
	}

	return &patient
}

func (adapter *PatientRepositoryAdapter) FindByPhoneNumber(phoneNumber string) *model.Patient {
	return &model.Patient{}
}

func (adapter *PatientRepositoryAdapter) FindByEmail(email string) *model.Patient {
	return &model.Patient{}
}

func (adapter *PatientRepositoryAdapter) Save(patient *model.Patient) *model.Patient {
	return &model.Patient{}
}
