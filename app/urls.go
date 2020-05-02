package app

import (
	"github.com/technolingo/bookstore-users-api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
