package users_repository

import (
	"context"
	"fmt"

	users_model "github.com/montinger-com/montinger-server/app/users/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.Background()

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) collection() *mongo.Collection {
	tenantMediaKey := fmt.Sprintf("%s", "users")
	collection := r.db.Collection(tenantMediaKey)
	return collection
}

func (r *UserRepository) GetByEmail(email string) (*users_model.User, error) {
	var user users_model.User
	filter := bson.M{"email": email, "status": bson.M{"$ne": "deleted"}}

	err := r.collection().FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
