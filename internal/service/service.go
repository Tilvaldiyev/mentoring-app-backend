package service

import "github.com/Tilvaldiyev/mentoring-app-backend/internal/model"

func (s *Service) GetCountries() ([]model.Country, error) {
	return s.repo.GetCountries()
}

func (s *Service) GetLanguages() ([]model.Language, error) {
	return s.repo.GetLanguages()
}

func (s *Service) GetExpertises() ([]model.Expertise, error) {
	return s.repo.GetExpertises()
}
