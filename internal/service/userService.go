package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"job-portal-api/internal/models"
	"job-portal-api/internal/pkg"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func (s *Service) UserSignIn(ctx context.Context, userData models.NewUser) (string, error) {

	// Check if the email exists in the user repository
	var userDetails models.User
	userDetails, err := s.UserRepo.CheckEmail(ctx, userData.Email)
	if err != nil {
		return "", err
	}

	// Check if the entered password matches the hashed password in the database
	err = pkg.CheckHashedPassword(userData.Password, userDetails.PasswordHash)
	if err != nil {
		log.Info().Err(err).Send()
		return "", errors.New("entered password is not wrong")
	}

	// Generate a JSON Web Token (JWT) for successful sign-in
	claims := jwt.RegisteredClaims{
		Issuer:    "job portal project",
		Subject:   strconv.FormatUint(uint64(userDetails.ID), 10),
		Audience:  jwt.ClaimStrings{"users"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token, err := s.auth.GenerateAuthToken(claims)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (s *Service) UserSignup(ctx context.Context, userData models.NewUser) (models.User, error) {
	hashedPass, err := pkg.HashPassword(userData.Password)
	if err != nil {
		return models.User{}, err
	}
	userDetails := models.User{
		Username:     userData.Username,
		Email:        userData.Email,
		PasswordHash: hashedPass,
	}
	userDetails, err = s.UserRepo.CreateUser(ctx, userDetails)
	if err != nil {
		return models.User{}, err
	}
	return userDetails, nil
}

func (s *Service) ForgotPassword(ctx context.Context, forgotPasswordDetails models.ForgotPasswordRequest) error {
	// Check if the email exists in the user repository
	userDetails, err := s.UserRepo.CheckEmail(ctx, forgotPasswordDetails.Email)
	if err != nil {
		return err
	}

	// Check if the provided username matches the username associated with the email
	if userDetails.Username != forgotPasswordDetails.Username {
		return errors.New("username does not match the email")
	}

	verficationCode := pkg.GenerateVerficationCode(forgotPasswordDetails.Email)

	err = s.rdb.VerficationCodeSet(ctx, userDetails.Email, verficationCode)
	if err != nil{
		return err
	}

	

	return nil
}

func(s *Service) PasswordRecovery(ctx context.Context,passwordRecoveryRequest models.PasswordRecoveryRequest)(error){
	// Check if the verfication code is valid
	verficationCode, err := s.rdb.VerficationCodeGet(ctx, passwordRecoveryRequest.Email)
	if err != nil{
		return err
	}
	if verficationCode!= passwordRecoveryRequest.VerficationCode{	
		return errors.New("verfication code is not valid")
	}
	// Update the password
	hashedPass, err := pkg.HashPassword(passwordRecoveryRequest.NewPassword)
	if err != nil {
		return err
	}
	err = s.UserRepo.UpdatePassword(ctx, passwordRecoveryRequest, hashedPass)
	if err != nil{
		return err
	}
	return nil
}
