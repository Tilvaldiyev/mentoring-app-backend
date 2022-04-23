package model

type MentorResponse struct {
	ID        int64  `json:"id"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Country   string `json:"country"`
	Language  string `json:"language"`
	Level     string `json:"level"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	Expertise string `json:"expertise"`
}
