package controllers

import (
	"Image-Processing-Service/database"
	"Image-Processing-Service/models"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UploadImage(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		log.Panic("Error while FormFile.")
		return
	}
	// Generate a unique filename using the current timestamp
	timestamp := time.Now().UnixNano()
	extension := filepath.Ext(file.Filename)
	uniqueFilename := fmt.Sprintf("%d%s", timestamp, extension)

	//upload file
	if err := c.SaveUploadedFile(file, "assets/uploads/"+uniqueFilename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save the file"})
		return
	}

	imageUrl := "http://127.0.0.1/uploads/" + uniqueFilename

	//image model
	var image = models.Image{
		ID:          primitive.NewObjectID(),
		Filename:    uniqueFilename,
		ContentType: file.Header.Get("Content-Type"),
		Size:        int16(file.Size),
		Url:         imageUrl,
	}

	//save image info to DB
	imgErr := database.SaveImageInfo(&image)
	if imgErr != nil {
		log.Print("Error While saving image info")
		return
	}

	//success
	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"data":    image,
	})

}

func ImageByID(c *gin.Context) {
	id := c.Param("id")

	var image *models.Image
	var err error
	image, err = database.GetImageInfoByID(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Image not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id"})
		return
	}

	//success
	c.JSON(http.StatusOK, gin.H{
		"data": image,
	})

}
