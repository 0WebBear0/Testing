package router

import (
	"Polytech-lyceum_back/view"
	"github.com/gin-gonic/gin"
)

func Router() {

	router := gin.Default()
	router.GET("", view.ShowAllUsersBack)

	//Users
	router.GET("/users", view.ShowAllUsers)
	router.POST("/user/add", view.CreateUser)

	router.Run()

}
