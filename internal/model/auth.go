package model

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PasswordRecoveryInput struct {
	Email      string `json:"email"`
	SecretCode string `json:"secret_code"`
}
