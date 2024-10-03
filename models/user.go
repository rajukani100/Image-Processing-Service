package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username" validate:"required"`
	Password  string             `json:"password" bson:"password" validate:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
