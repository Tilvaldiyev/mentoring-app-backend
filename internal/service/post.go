package service

import "github.com/Tilvaldiyev/mentoring-app-backend/internal/model"

func (s *Service) CreatePost(payload model.Posts) error {
	return s.repo.CreatePost(payload)
}

func (s *Service) GetPost(id int) (model.Posts, error) {
	return s.repo.GetPost(id)
}

func (s *Service) DeletePost(id int) error {
	return s.repo.DeletePost(id)
}
