package database

import (
	"Image-Processing-Service/models"
	"context"
)

func SaveImageInfo(image *models.Image) error {
	collection := client.Database("ImageProcessingService").Collection("images")
	_, err := collection.InsertOne(context.Background(), *image)
	return err
}
