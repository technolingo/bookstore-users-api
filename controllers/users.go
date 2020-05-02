package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/technolingo/bookstore-users-api/models"
	"github.com/technolingo/bookstore-users-api/services"
)

// CreateUser creates a user object according to the payload
func CreateUser(c *gin.Context) {
	var user models.User

	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	return
	// }
	if err := c.ShouldBindJSON(&user); err != nil {
		// TODO: handle error
		return
	}

	result, err := services.CreateUser(&user)
	if err != nil {
		// TODO: handle error
		return
	}

	c.JSON(http.StatusCreated, result)
}

// RetrieveUser gets a user object according to the `user_id`
func RetrieveUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

// SearchUser returns a list of users matching the query param supplied
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

// UpdateUser updates an existing user according to the payload
func UpdateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

// DeleteUser removes the user object according to the `user_id`
func DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
