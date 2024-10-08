package monitors_repository

import (
	"context"
	"fmt"

	monitors_model "github.com/montinger-com/montinger-server/app/monitors/models"
	"github.com/montinger-com/montinger-server/app/utils/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *MonitorsRepository) Create(monitor *monitors_model.Monitor) error {
	collection := r.collection()

	created, err := collection.InsertOne(ctx, monitor)
	if err != nil {
		return err
	}
	monitor.ID = helpers.ObjectIDToString(created.InsertedID)

	return nil
}

func (r *MonitorsRepository) GetByID(id string) (*monitors_model.Monitor, error) {
	collection := r.collection()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var monitor monitors_model.Monitor
	err = collection.FindOne(ctx, bson.M{"_id": objectId, "status": bson.M{"$eq": "active"}}).Decode(&monitor)
	if err != nil {
		return nil, err
	}

	return &monitor, nil
}

func (r *MonitorsRepository) Update(monitor *monitors_model.Monitor) error {
	collection := r.collection()

	_, err := collection.UpdateOne(ctx, bson.M{"_id": monitor.ID, "status": bson.M{"$eq": "active"}}, bson.M{"$set": monitor})
	if err != nil {
		return err
	}

	return nil
}
