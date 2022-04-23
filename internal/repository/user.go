package repository

import (
	"fmt"
	"github.com/Tilvaldiyev/mentoring-app-backend/internal/model"
)

func (r *Repository) GetUserTypes() ([]model.UserType, error) {
	var userTypes []model.UserType

	err := r.DB.Find(&userTypes).Error
	if err != nil {
		return userTypes, fmt.Errorf("getting user types err: %w", err)
	}

	return userTypes, nil
}

func (r *Repository) GetUserLevels() ([]model.Level, error) {
	var userLevels []model.Level

	err := r.DB.Find(&userLevels).Error
	if err != nil {
		return userLevels, fmt.Errorf("getting user levels err: %w", err)
	}

	return userLevels, nil
}
