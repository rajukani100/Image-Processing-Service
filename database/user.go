package database

import (
	"Image-Processing-Service/models"
	"context"
)

func SaveUserToDB(user *models.User) error {
	collection := client.Database("ImageProcessingService").Collection("users")
	_, err := collection.InsertOne(context.Background(), *user)
	return err
}
