package repository

import "github.com/Tilvaldiyev/mentoring-app-backend/internal/model"

type RepositoryInt interface {
	AuthInt
	ArticleInt
	PostInt
	UserInt
	MentorInt
	GetCountries() ([]model.Country, error)
	GetLanguages() ([]model.Language, error)
	GetExpertises() ([]model.Expertise, error)
}

type AuthInt interface {
	CreateUser(user model.Users) error
	GetUser(email, password string) (model.Users, error)
	UserExistence(email string) (string, error)
	VerifySecretCode(email, secretCode string) error
	UpdatePassword(email, password, newSecretCode string) error
}

type ArticleInt interface {
	CreateArticle(payload model.Article) error
	GetArticle(id int) (model.Article, error)
	DeleteArticle(id int) error
	UpdateArticle(id int, payload model.UpdateArticleInput) error
}

type PostInt interface {
	CreatePost(payload model.Posts) error
	GetPost(id int) (model.Posts, error)
	DeletePost(id int) error
}

type UserInt interface {
	GetUserTypes() ([]model.UserType, error)
	GetUserLevels() ([]model.Level, error)
}

type MentorInt interface {
	GetMentorsList(searchMap map[string]string) ([]model.MentorResponse, error)
}
