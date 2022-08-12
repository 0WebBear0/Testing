package view

import (
	"Polytech-lyceum_back/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Date struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func ShowAllUsersBack(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	context.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
	var date = []Date{
		{Name: "1", Description: "2"},
		{Name: "3", Description: "4"},
		{Name: "5", Description: "6"},
		{Name: "7", Description: "8"},
	}

	context.JSON(http.StatusOK,
		gin.H{
			"date": date,
		})
}

func ShowAllUsers(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	context.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
	var users []models.Test

	models.DB().Find(&users)

	context.JSON(http.StatusOK, gin.H{"users": users})
}

type CreateUserInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func CreateUser(context *gin.Context) {
	var userInput CreateUserInput

	if err := context.ShouldBindJSON(&userInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + " Неправильно введены данные"})
		return
	}

	user := models.Test{
		Name:        userInput.Name,
		Description: userInput.Description,
	}

	models.DB().Create(&user)

	context.JSON(http.StatusOK, gin.H{"user": user})

}
