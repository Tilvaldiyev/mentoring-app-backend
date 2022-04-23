package repository

import (
	"fmt"
	"github.com/Tilvaldiyev/mentoring-app-backend/internal/model"
)

func (r *Repository) CreatePost(post model.Posts) error {
	err := r.DB.Create(&post).Error
	if err != nil {
		return fmt.Errorf("creating post err: %w", err)
	}

	return nil
}

func (r *Repository) GetPost(id int) (model.Posts, error) {
	var post model.Posts

	err := r.DB.Where("id=?", id).First(&post).Error
	if err != nil {
		return post, fmt.Errorf("getting post err: %w", err)
	}

	return post, nil
}

func (r *Repository) DeletePost(id int) error {
	err := r.DB.Delete(&model.Posts{}, id).Error
	if err != nil {
		return fmt.Errorf("deleting post err: %w", err)
	}

	return nil
}
