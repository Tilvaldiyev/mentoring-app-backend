package repository

import (
	"errors"
	"fmt"
	"github.com/Tilvaldiyev/mentoring-app-backend/internal/model"
	"gorm.io/gorm"
)

func (r *Repository) CreateUser(user model.Users) error {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := r.DB.Create(&user).Error
		if err != nil {
			return fmt.Errorf("creating user err: %w", err)
		}

		var expertise model.ExpertiseUser
		expertise.UserID = user.ID
		for _, exID := range user.ExpertiseIDs {
			expertise.ExpertiseID = exID
			err := r.DB.Create(&expertise).Error
			if err != nil {
				return fmt.Errorf("creating expertise err: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("transaction err: %w", err)
	}

	return nil
}

func (r *Repository) GetUser(email, password string) (model.Users, error) {
	var user model.Users

	err := r.DB.Where("email=? and password=?", email, password).First(&user).Error
	if err != nil {
		return user, fmt.Errorf("getting user err: %w", err)
	}

	return user, nil
}

func (r *Repository) UserExistence(email string) (string, error) {
	var user model.Users

	err := r.DB.Where("email=?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("user with email %s not exist", email)
		}
		return "", fmt.Errorf("getting user err: %w", err)
	}

	return user.SecretCode, nil
}

func (r *Repository) VerifySecretCode(email, secretCode string) error {
	var user model.Users

	err := r.DB.Where("email=? and secret_code=?", email, secretCode).First(&user).Error
	if err != nil {
		return fmt.Errorf("verifying secret code err: %w", err)
	}

	return nil
}

func (r *Repository) UpdatePassword(email, password, newSecretCode string) error {
	err := r.DB.Model(model.Users{}).Where("email=?", email).Updates(map[string]interface{}{"password": password, "secret_code": newSecretCode}).Error
	if err != nil {
		return fmt.Errorf("update password err: %w", err)
	}

	return nil
}
