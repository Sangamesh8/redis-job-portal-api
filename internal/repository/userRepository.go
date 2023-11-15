package repository

import (
	"context"
	"errors"

	"job-portal-api/internal/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateUser(ctx context.Context, UserDetails models.User) (models.User, error) {
	// Use the database instance to create a new user
	result := r.DB.Create(&UserDetails)
	// Check for errors during the database insertion
	if result.Error != nil {
		// Log the error and return an error indicating that user creation failed
		log.Info().Err(result.Error).Send()
		return models.User{}, errors.New("could not create the user")
	}
	//return the created user details
	return UserDetails, nil
}

func (r *Repo) CheckEmail(ctx context.Context, email string) (models.User, error) {
	// Declare a variable to store user details
	var userDetails models.User
	// Query the database to find a user with the specified email
	result := r.DB.Where("email = ?", email).First(&userDetails)
	if result.Error != nil {
		// Log the error and return an error indicating that the email is not found
		log.Info().Err(result.Error).Send()
		return models.User{}, errors.New("email not found")
	}
	// If no error occurred, return the retrieved user details
	return userDetails, nil
}
