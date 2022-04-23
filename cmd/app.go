package cmd

import (
	"context"
	"github.com/Tilvaldiyev/mentoring-app-backend/config"
	"github.com/Tilvaldiyev/mentoring-app-backend/internal/handler"
	"github.com/Tilvaldiyev/mentoring-app-backend/internal/repository"
	"github.com/Tilvaldiyev/mentoring-app-backend/internal/service"
	"github.com/Tilvaldiyev/mentoring-app-backend/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func Run() {
	logger, err := logging.Init()
	if err != nil {
		panic(err)
	}

	logger.Info("loading environment variables...")
	err = config.LoadEnv()
	if err != nil {
		logger.Errorf("error occured while loading environment variables: %s", err.Error())
		panic(err)
	}
	logger.Info("environment variables successfully loaded")

	logger.Info("initializing configuration...")
	cfg, err := config.InitConfig("config.json")
	if err != nil {
		logger.Errorf("error occured while initializing configuration: %s", err.Error())
		panic(err)
	}
	logger.Info("configuration successfully initialized")

	logger.Info("connecting to DB")
	db, err := repository.NewDB(cfg.DB)
	if err != nil {
		logger.Errorf("error occured while connecting to DB: %s", err.Error())
		panic(err)
	}
	logger.Info("successfully connected to DB")

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(cfg, logger, services)
	routers := handlers.InitRouter()

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(cfg.Server.Port),
		Handler: routers,
	}

	if cfg.Server.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	go func() {
		logger.Infof("running server on port %s...", strconv.Itoa(cfg.Server.Port))
		err = server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Errorf("error occured while running server: %s", err.Error())
			panic(err)
		}
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm

	srvctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	logger.Info("shutting down server...")
	err = server.Shutdown(srvctx)
	if err != nil {
		logger.Errorf("error occured while shutting down server: %s", err.Error())
	}
}
