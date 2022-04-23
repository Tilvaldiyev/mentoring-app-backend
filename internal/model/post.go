package model

import "time"

type Posts struct {
	ID          int64     `json:"id" gorm:"not null;primary_key:true"`
	Description string    `json:"description" gorm:"not null;"`
	UserID      int64     `json:"user_id"`
	User        Users     `json:"-"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:now()"`
}

func (e Posts) TableName() string {
	return "posts"
}
