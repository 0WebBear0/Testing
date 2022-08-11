package router

import (
	"Polytech-lyceum_back/view"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() {

	router := gin.Default()

	router.GET("", func(context *gin.Context) {
		context.String(http.StatusOK, "World")
	})

	//Users
	router.GET("/users", view.ShowAllUsers)
	router.POST("/user/add", view.CreateUser)

	router.Run()

}
