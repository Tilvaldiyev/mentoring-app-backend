package service

import "github.com/Tilvaldiyev/mentoring-app-backend/internal/model"

type ServiceInt interface {
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
	CreateUser(payload model.Users) error
	GenerateToken(payload model.SignInInput) (string, error)
	ParseToken(token string) (int64, error)
	VerifyEmailAndSendCode(payload model.PasswordRecoveryInput) (string, error)
	VerifySecretCode(payload model.PasswordRecoveryInput) error
	RecoverPassword(payload model.SignInInput) error
}

type ArticleInt interface {
	CreateArticle(payload model.Article) error
	GetArticle(id int) (model.Article, error)
	DeleteArticle(id int) error
	UpdateArticle(idInt int, req model.UpdateArticleInput) error
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
