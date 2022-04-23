package repository

import (
	"fmt"
	"github.com/Tilvaldiyev/mentoring-app-backend/internal/model"
)

func(r *Repository) GetCountries() ([]model.Country, error) {
	var countries []model.Country

	err := r.DB.Find(&countries).Error
	if err != nil {
		return countries, fmt.Errorf("getting countries err: %w", err)
	}

	return countries, nil
}

func(r *Repository) GetLanguages() ([]model.Language, error) {
	var languages []model.Language

	err := r.DB.Find(&languages).Error
	if err != nil {
		return languages, fmt.Errorf("getting languages err: %w", err)
	}

	return languages, nil
}

func(r *Repository) GetExpertises() ([]model.Expertise, error) {
	var expertise []model.Expertise

	err := r.DB.Find(&expertise).Error
	if err != nil {
		return expertise, fmt.Errorf("getting expertise err: %w", err)
	}

	return expertise, nil
}
