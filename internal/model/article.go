package model

import "time"

type Article struct {
	ID           int64     `json:"id" gorm:"not null;primary_key:true"`
	Title        string    `json:"title" gorm:"not null;size:255;"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:now()"`
	UserID       int64     `json:"user_id"`
	User         Users     `json:"-"`
	ImgURL       string    `json:"img_url"`
	Views        int       `json:"views" gorm:"0"`
	Likes        int       `json:"likes" gorm:"0"`
	LevelID      int64     `json:"level_id"`
	Level        Level     `json:"-"`
	ExpertiseIDs []int64   `json:"expertise_ids" gorm:"-"`
}

func (e Article) TableName() string {
	return "article"
}

type UpdateArticleInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}
