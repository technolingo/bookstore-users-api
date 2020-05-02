package services

import (
	"github.com/technolingo/bookstore-users-api/errors"
	"github.com/technolingo/bookstore-users-api/models"
)

// CreateUser processes the buisness logic and creates the user
func CreateUser(u *models.User) (*models.User, *errors.APIError) {
	return &models.User{}, nil
}
