package repository

import (
	"fmt"
	"github.com/Tilvaldiyev/mentoring-app-backend/internal/model"
	"gorm.io/gorm"
)

func (r *Repository) CreateArticle(article model.Article) error {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := r.DB.Create(&article).Error
		if err != nil {
			return fmt.Errorf("creating article err: %w", err)
		}

		var expertise model.ExpertiseArticle
		expertise.ArticleID = article.ID
		for _, exID := range article.ExpertiseIDs {
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

func (r *Repository) GetArticle(id int) (model.Article, error) {
	var article model.Article

	err := r.DB.Where("id=?", id).First(&article).Error
	if err != nil {
		return article, fmt.Errorf("getting article err: %w", err)
	}

	return article, nil
}

func (r *Repository) DeleteArticle(id int) error {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := r.DB.Where("article_id=?", id).Delete(&model.ExpertiseArticle{}).Error
		if err != nil {
			return fmt.Errorf("deleting expertise article err: %w", err)
		}

		err = r.DB.Delete(&model.Article{}, id).Error
		if err != nil {
			return fmt.Errorf("deleting article err: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("transaction err: %w", err)
	}

	return nil
}

func (r *Repository) UpdateArticle(id int, payload model.UpdateArticleInput) error {
	updateMap := make(map[string]interface{})
	if payload.Title != nil {
		updateMap["title"] = payload.Title
	}
	if payload.Description != nil {
		updateMap["description"] = payload.Description
	}

	err := r.DB.Model(model.Article{}).Where("id=?", id).Updates(updateMap).Error
	if err != nil {
		return fmt.Errorf("update err: %w", err)
	}

	return nil

}
