package database

import (
	"Image-Processing-Service/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveUserToDB(user *models.User) error {
	collection := client.Database("ImageProcessingService").Collection("users")
	_, err := collection.InsertOne(context.Background(), *user)
	return err
}

func CheckUserExist(username string) (bool, error) {
	var user models.User
	collection := client.Database("ImageProcessingService").Collection("users")
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil // User does not exist
		}
		return false, err // Return the error for other cases
	}
	return true, nil // User exists
}
