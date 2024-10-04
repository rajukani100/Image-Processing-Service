package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Image struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	ContentType string             `json:"content_type" bson:"content_type"`
	Filename    string             `json:"filename" bson:"filename"`
	Size        int16              `json:"size" bson:"size"`
	Url         string             `json:"url" bson:"url"`
}
