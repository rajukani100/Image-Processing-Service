package database

import (
	"Image-Processing-Service/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetImageList(page int, limit int) (*[]models.Image, error) {
	collection := client.Database("ImageProcessingService").Collection("images")

	skip := int64((page - 1) * limit)
	opts := options.Find().SetSkip(skip).SetLimit(int64(limit))

	cur, err := collection.Find(context.Background(), bson.M{"url": bson.M{"$regex": "/uploads/"}}, opts)
	if err != nil {
		return nil, err
	}
	//images list
	var results []models.Image
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return &results, nil
}
