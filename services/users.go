package services

import (
	"github.com/technolingo/bookstore-users-api/models"
	"github.com/technolingo/bookstore-users-api/utils"
)

// CreateUser processes the buisness logic and creates the user
func CreateUser(u *models.User) (*models.User, *utils.APIError) {
	return &models.User{}, nil
}
