package database

import (
	"Image-Processing-Service/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveImageInfo(image *models.Image) error {
	collection := client.Database("ImageProcessingService").Collection("images")
	_, err := collection.InsertOne(context.Background(), *image)
	return err
}

func GetImageInfoByID(id string) (*models.Image, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := client.Database("ImageProcessingService").Collection("images")

	var image models.Image
	err = collection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&image)
	if err != nil {
		return nil, err
	}
	return &image, nil

}
