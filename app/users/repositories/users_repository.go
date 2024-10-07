package users_repository

import (
	"context"
	"fmt"

	users_model "github.com/montinger-com/montinger-server/app/users/models"
	"github.com/montinger-com/montinger-server/app/utils/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.Background()

type UsersRepository struct {
	db *mongo.Database
}

func NewUsersRepository(db *mongo.Database) *UsersRepository {
	return &UsersRepository{db}
}

func (r *UsersRepository) collection() *mongo.Collection {
	collectionName := fmt.Sprintf("%s", "users")
	collection := r.db.Collection(collectionName)
	return collection
}

func (r *UsersRepository) GetByEmail(email string) (*users_model.User, error) {
	var user users_model.User
	filter := bson.M{"email": email, "status": bson.M{"$ne": "deleted"}}

	err := r.collection().FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UsersRepository) Create(user *users_model.User) error {
	collection := r.collection()
	created, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	user.ID = helpers.ObjectIDToString(created.InsertedID)

	return nil
}
