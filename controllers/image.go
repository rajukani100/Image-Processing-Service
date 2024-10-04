package controllers

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
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

	//success
	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"data": gin.H{
			"filename":    uniqueFilename,
			"url":         imageUrl,
			"size":        file.Size,
			"contentType": file.Header.Get("Content-Type"),
		},
	})

}
