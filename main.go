package main

import (
	"github.com/Tilvaldiyev/mentoring-app-backend/cmd"
)

// @title Mentoring App API
// @version 1.0
// @description API for virtual mentoring application
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cmd.Run()
}
