package service

import (
	"crypto/tls"
	"fmt"
	"github.com/Tilvaldiyev/mentoring-app-backend/internal/model"
	"github.com/Tilvaldiyev/mentoring-app-backend/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/gomail.v2"
	"os"
	"time"
)

const (
	tokenTTL   = 12 * time.Hour
	signingKey = "fsdfdsfsdfsdfsd"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int64 `json:"user_id"`
}

func (s *Service) CreateUser(payload model.Users) error {
	payload.Password = utils.GeneratePassword(payload.Password)
	payload.SecretCode = utils.GenerateRandomSecretCode()
	return s.repo.CreateUser(payload)
}

func (s *Service) GenerateToken(payload model.SignInInput) (string, error) {
	user, err := s.repo.GetUser(payload.Email, utils.GeneratePassword(payload.Password))
	if err != nil {
		return "", fmt.Errorf("DB err: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *Service) VerifyEmailAndSendCode(payload model.PasswordRecoveryInput) (string, error) {
	secretCode, err := s.repo.UserExistence(payload.Email)
	if err != nil {
		return "", fmt.Errorf("DB err: %w", err)
	}

	err = sendCodeToEmail(payload.Email, secretCode)
	if err != nil {
		return "", fmt.Errorf("send code to email err: %w", err)
	}

	return "code sent to your email", nil
}

func (s *Service) RecoverPassword(payload model.SignInInput) error {
	newSecretCode := utils.GenerateRandomSecretCode()
	err := s.repo.UpdatePassword(payload.Email, utils.GeneratePassword(payload.Password), newSecretCode)
	if err != nil {
		return fmt.Errorf("update password err: %w", err)
	}

	return nil
}

func (s *Service) VerifySecretCode(payload model.PasswordRecoveryInput) error {
	err := s.repo.VerifySecretCode(payload.Email, payload.SecretCode)
	if err != nil {
		return fmt.Errorf("update password err: %w", err)
	}

	return nil
}

func (s *Service) ParseToken(accessToken string) (int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, fmt.Errorf("parse token err: %w", err)
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, fmt.Errorf("token claims are not of type tokenClaims")
	}

	return claims.UserID, nil
}

func sendCodeToEmail(email, secretCode string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", os.Getenv("EMAIL"))
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Password Recovery Code")
	message.SetBody("text/plain", "This is your secret code to recover password: "+secretCode)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("EMAIL_PASSWORD"))
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := dialer.DialAndSend(message)
	if err != nil {
		return fmt.Errorf("DialAndSend err: %w", err)
	}

	return nil
}
