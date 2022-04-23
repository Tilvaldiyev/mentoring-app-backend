package handler

import (
	"github.com/Tilvaldiyev/mentoring-app-backend/config"
	"github.com/Tilvaldiyev/mentoring-app-backend/internal/service"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	config   *config.Configuration
	log      *logrus.Entry
	services service.ServiceInt
}

func NewHandler(cfg *config.Configuration, log *logrus.Entry, services service.ServiceInt) *Handler {
	return &Handler{config: cfg, log: log, services: services}
}
