package mongodb

import (
	"context"
	"time"

	"github.com/chava.gnolasco/polaris/infraestructure/env"
	"github.com/chava.gnolasco/polaris/infraestructure/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient mongo.Client

/**
 * GetCollection
 * @param collectionName string
 * @return *mongo.Collection
 */
func GetCollection(collectionName string) *mongo.Collection {
	return mongoClient.Database(env.GetEnv().MONGO_DB_NAME).Collection(collectionName)
}

/**
 * CloseConnection
 *
 * It closes the connection to the MongoDB database
 */
func CloseConnection() {
	log.Info("[MongoDb] Closing the connection...")
	if err := mongoClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

/**
 * MongoDbInit
 *
 * It initializes the connection to the MongoDB database
 */
func init() {
	log.Info("[MongoDb] Initializing the connection...")
	uri := env.GetEnv().MONGO_URL
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	mongoClient = *client
}
