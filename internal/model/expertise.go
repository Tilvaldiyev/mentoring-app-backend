package model

type Expertise struct {
	ID   int64  `json:"id" gorm:"not null;primary_key:true"`
	Name string `json:"name" gorm:"not null; size:255"`
}

func (e Expertise) TableName() string {
	return "expertise"
}

type ExpertiseUser struct {
	ID          int64     `json:"id" gorm:"not null;primary_key:true"`
	UserID      int64     `json:"user_id"`
	User        Users     `json:"-"`
	ExpertiseID int64     `json:"expertise_id"`
	Expertise   Expertise `json:"-"`
}

func (e ExpertiseUser) TableName() string {
	return "expertise_user"
}

type ExpertiseArticle struct {
	ID          int64     `json:"id" gorm:"not null;primary_key:true"`
	ArticleID   int64     `json:"article_id"`
	Article     Article   `json:"-"`
	ExpertiseID int64     `json:"expertise_id"`
	Expertise   Expertise `json:"-"`
}

func (e ExpertiseArticle) TableName() string {
	return "expertise_article"
}
