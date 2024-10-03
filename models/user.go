package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Email        string             `json:"email" bson:"email" validate:"required,email"`
	HashPassword string             `json:"hash_password" bson:"hash_password" validate:"required"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
}
