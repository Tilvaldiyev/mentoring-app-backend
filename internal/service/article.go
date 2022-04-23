package service

import "github.com/Tilvaldiyev/mentoring-app-backend/internal/model"

func (s *Service) CreateArticle(payload model.Article) error {
	return s.repo.CreateArticle(payload)
}

func (s *Service) GetArticle(id int) (model.Article, error) {
	return s.repo.GetArticle(id)
}

func (s *Service) DeleteArticle(id int) error {
	return s.repo.DeleteArticle(id)
}

func (s *Service) UpdateArticle(idInt int, req model.UpdateArticleInput) error {
	return s.repo.UpdateArticle(idInt, req)
}
