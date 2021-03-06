package app

import (
	"github.com/gin-gonic/gin"

	"github.com/technolingo/bookstore-users-api/controllers"
)

func mapUrls(router *gin.Engine) {
	router.GET("/ping", controllers.Ping)

	router.GET("/users/:user_id", controllers.RetrieveUser)
	router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", controllers.CreateUser)
	router.PATCH("/users/:user_id", controllers.UpdateUser)
	router.DELETE("/users/:user_id", controllers.DeleteUser)
}
