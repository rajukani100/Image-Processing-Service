package controllers

import (
	"Image-Processing-Service/database"
	"Image-Processing-Service/models"
	"Image-Processing-Service/services"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Register(c *gin.Context) {
	var user models.User
	if bindErr := c.BindJSON(&user); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while binding"})
		return
	}

	//custom validation
	validate := validator.New()
	validateErr := validate.Struct(user)
	if validateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "field must not be empty."})
		return
	}

	//check if user already exist
	isUserExist, userExistErr := database.CheckUserExist(user.Username)
	if userExistErr != nil {
		log.Panic("Error While Finding User.")
		return
	}

	if isUserExist {
		c.JSON(http.StatusConflict, gin.H{"error": "User already Exist."})
		return
	}

	//hashing password
	hashPassword, err := services.HashPassword(user.Password)
	if err != nil {
		log.Panic("Error while hashing password")
		return
	}
	//initializing user
	user.ID = primitive.NewObjectID()
	user.Password = hashPassword
	user.CreatedAt = time.Now()

	//save User to DB
	err = database.SaveUserToDB(&user)
	if err != nil {
		log.Panic("Error While saving user to DB.")
		return
	}

	//generating JWT
	jwtToken, tokenErr := services.GenerateJwt(&user.Username)
	if tokenErr != nil {
		log.Panic("Error While JWT token Generation.")
		return
	}

	//success response
	response := struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}{User: user, Token: jwtToken}

	c.JSON(http.StatusOK, response)

}
