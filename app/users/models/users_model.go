package users_model

import "time"

type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password,omitempty" bson:"password"`
	Status   string `json:"status" bson:"status"`

	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deleted_at"`
}
