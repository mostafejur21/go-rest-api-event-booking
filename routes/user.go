package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mostafejur21/go_rest_api/models"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user data"})
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user data"})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
