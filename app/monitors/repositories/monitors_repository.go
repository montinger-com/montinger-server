package monitors_repository

import (
	"context"
	"fmt"

	monitors_model "github.com/montinger-com/montinger-server/app/monitors/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.Background()

type MonitorsRepository struct {
	db *mongo.Database
}

func NewMonitorsRepository(db *mongo.Database) *MonitorsRepository {
	return &MonitorsRepository{db}
}

func (r *MonitorsRepository) collection() *mongo.Collection {
	tenantMediaKey := fmt.Sprintf("%s", "monitors")
	collection := r.db.Collection(tenantMediaKey)
	return collection
}

func (r *MonitorsRepository) GetAll() ([]*monitors_model.Monitor, error) {
	collection := r.collection()

	cursor, err := collection.Find(ctx, bson.M{"status": bson.M{"$eq": "active"}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var monitors []*monitors_model.Monitor

	if err = cursor.All(ctx, &monitors); err != nil {
		return nil, err
	}

	return monitors, nil
}
