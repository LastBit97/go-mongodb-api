package services

import "github.com/LastBit97/go-mongodb-api/models"

type AuthService interface {
	SignUpUser(*models.SignUpInput) (*models.DBResponse, error)
	VerifyEmail(verificationCode string) error
	UpdatePasswordResetTokenByEmail(email string, passwordResetToken string) error
	ResetPassword(resetToken string, password string) error
}
