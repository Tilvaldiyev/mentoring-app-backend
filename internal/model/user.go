package model

type Users struct {
	ID           int64    `json:"id" gorm:"not null;primary_key:true"`
	LastName     string   `json:"last_name" gorm:"not null; size:255;"`
	FirstName    string   `json:"first_name" gorm:"not null;size:255;"`
	Email        string   `json:"email" gorm:"not null;size:255;"`
	Password     string   `json:"password" gorm:"not null;size:255;"`
	UserTypeID   int64    `json:"user_type_id"`
	CountryID    int64    `json:"country_id"`
	Country      Country  `json:"-"`
	LanguageID   int64    `json:"language_id"`
	Language     Language `json:"-"`
	LevelID      int64    `json:"level_id"`
	Level        Level    `json:"-"`
	Title        string   `json:"title" gorm:"not null;size:255;"`
	Info         string   `json:"info"`
	SecretCode   string   `json:"-" gorm:"size:255"`
	UserType     UserType `json:"-"`
	ExpertiseIDs []int64  `json:"expertise_ids" gorm:"-"`
}

func (u Users) TableName() string {
	return "users"
}

type Language struct {
	ID   int64  `json:"id" gorm:"not null;primary_key:true"`
	Name string `json:"name" gorm:"not null; size:255"`
}

func (l Language) TableName() string {
	return "language"
}

type Country struct {
	ID   int64  `json:"id" gorm:"not null;primary_key:true"`
	Name string `json:"name" gorm:"not null; size:255"`
}

func (c Country) TableName() string {
	return "country"
}

type UserType struct {
	ID   int64  `json:"id" gorm:"not null;primary_key:true"`
	Name string `json:"name" gorm:"not null; size:255"`
}

func (e UserType) TableName() string {
	return "user_type"
}
