package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApp is the entrypoint of the application
func StartApp() {
	mapUrls()
	router.Run(":8080")
}
