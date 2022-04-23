package service

import "github.com/Tilvaldiyev/mentoring-app-backend/internal/model"

func (s *Service) GetUserTypes() ([]model.UserType, error) {
	return s.repo.GetUserTypes()
}

func (s *Service) GetUserLevels() ([]model.Level, error) {
	return s.repo.GetUserLevels()
}
