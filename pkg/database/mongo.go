package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	db *mongo.Database
}

func NewMongoClient(ctx context.Context, host string, port string, username string, password string, dbName string) (*MongoClient, error) {
	connectionString := fmt.Sprintf("mongodb://%v:%v@%v:%v/?authSource=%v", username, password, host, port, dbName)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to create MongoDB connection: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	return &MongoClient{db: client.Database(dbName)}, nil
}

func (mongo *MongoClient) GetDatabaseDetails() string {

	var commandResult bson.M
	command := bson.D{{Key: "serverStatus", Value: 1}}
	err := mongo.db.RunCommand(context.TODO(), command).Decode(&commandResult)

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%v", commandResult["version"])

}

func (mongo *MongoClient) GetDB() *mongo.Database {
	return mongo.db
}
