package db

import (
	"context"
	"database/sql"

	"github.com/montinger-com/montinger-server/config"
	"github.com/montinger-com/montinger-server/pkg/database"
	"github.com/rashintha/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

var rawClient any

var MongoClient *mongo.Database
var SqlClient *sql.DB

func init() {
	logger.Defaultln("Initializing Database")

	if config.DB_TYPE == database.MONGODB {

		client, err := database.NewMongoClient(
			context.Background(),
			config.DB_HOST,
			config.DB_PORT,
			config.DB_USER,
			config.DB_PASS,
			config.DB_NAME,
		)
		if err != nil {
			logger.Errorf("Failed to initialize database: %v\n", err.Error())
		}

		if err == nil {
			logger.Defaultln("Mongo database initialized")
		}
		rawClient = client
		MongoClient = client.GetDB()

	} else {

		client, err := database.NewSqlClient(
			context.Background(),
			config.DB_TYPE,
			config.DB_HOST,
			config.DB_PORT,
			config.DB_USER,
			config.DB_PASS,
			config.DB_NAME,
		)

		if err != nil {
			logger.Errorln("Failed to initialize database")
		}

		if err == nil {
			logger.Defaultln("SQL database initialized")
		}
		rawClient = client
		SqlClient = client.GetDB()

	}
}
