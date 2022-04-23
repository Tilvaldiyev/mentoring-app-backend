package service

import "github.com/Tilvaldiyev/mentoring-app-backend/internal/repository"

type Service struct {
	repo repository.RepositoryInt
}

func NewService(repo repository.RepositoryInt) *Service {
	return &Service{repo: repo}
}

