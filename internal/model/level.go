package model

type Level struct {
	ID   int64  `json:"id" gorm:"not null;primary_key:true"`
	Name string `json:"name" gorm:"not null; size:255"`
}

func (e Level) TableName() string {
	return "level"
}

