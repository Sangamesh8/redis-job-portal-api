package models

import "gorm.io/gorm"

type NewUser struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	gorm.Model
	Username     string `json:"username" gorm:"unique"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash string `json:"-"`
}

type ForgotPasswordRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type PasswordRecoveryRequest struct {
	VerficationCode    string `json:"verficationCode" validate:"required"`
	Email              string `json:"email" validate:"required"`
	NewPassword        string `json:"newPassword" validate:"required"`
	NewPasswordConfirm string `json:"newPasswordConfirm" validate:"required"`
}
