package controllers

import (
	"Image-Processing-Service/database"
	"Image-Processing-Service/models"
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
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
	uploadImgPath := "assets/uploads/" + uniqueFilename
	if err := c.SaveUploadedFile(file, uploadImgPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save the file"})
		return
	}

	imageUrl := "http://127.0.0.1/uploads/" + uniqueFilename

	//size of new image
	imgSize, err := os.Stat(uploadImgPath)
	if err != nil {
		log.Print("Error while calculating size")
		return
	}

	//image model
	var image = models.Image{
		ID:          primitive.NewObjectID(),
		Filename:    uniqueFilename,
		ContentType: file.Header.Get("Content-Type"),
		Size:        int16(imgSize.Size()),
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

func TransformImage(c *gin.Context) {
	id := c.Param("id")

	//original img
	var img *models.Image
	var err error
	img, err = database.GetImageInfoByID(id)
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

	// transform model
	var transformRequest models.TransformRequest
	if bindErr := c.BindJSON(&transformRequest); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while binding"})
		return
	}

	imgPath := "assets/uploads/" + img.Filename
	imgFile, err := imgio.Open(imgPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while opening image"})
		return
	}

	if transformRequest.Transformations.Resize != nil {
		//resize
		imgFile = transform.Resize(imgFile, transformRequest.Transformations.Resize.Width, transformRequest.Transformations.Resize.Height, transform.Linear)
	}

	if transformRequest.Transformations.Crop != nil {
		//crop
		c := transformRequest.Transformations.Crop
		imgFile = transform.Crop(imgFile, image.Rect(c.X0, c.Y0, c.X1, c.Y1))
	}

	if transformRequest.Transformations.Rotate != nil {
		//rotate
		imgFile = transform.Rotate(imgFile, *transformRequest.Transformations.Rotate, nil)
	}

	if transformRequest.Transformations.Filters != nil {
		//filters
		if transformRequest.Transformations.Filters.Grayscale {
			imgFile = effect.Grayscale(imgFile)
		}
		if transformRequest.Transformations.Filters.Sepia {
			imgFile = effect.Sepia(imgFile)
		}
	}

	transformedImgPath := "assets/edited/" + img.Filename
	if err := imgio.Save(transformedImgPath, imgFile, imgio.PNGEncoder()); err != nil {
		log.Print("Error while saving transformed output.")
		return
	}

	//size of new image
	imgSize, err := os.Stat(transformedImgPath)
	if err != nil {
		log.Print("Error while calculating size")
		return
	}

	//transformed image model
	var transformedImg = models.Image{
		ID:          primitive.NewObjectID(),
		ContentType: img.ContentType,
		Size:        int16(imgSize.Size()),
		Filename:    img.Filename,
		Url:         "http://127.0.0.1/edited/" + img.Filename,
	}

	//saving transformed image data to DB
	err = database.SaveImageInfo(&transformedImg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while saving image info"})
		return
	}

	//success
	c.JSON(http.StatusOK, gin.H{"data": transformedImg})
}
