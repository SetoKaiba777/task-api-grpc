package loader

import (
	"context"
	"go-challenger/adapter/repository"
	coreRepository "go-challenger/core/repository"
	"go-challenger/core/usecase"
	appConfig "go-challenger/infrastructure/config"
	"go-challenger/infrastructure/database"
	"go-challenger/infrastructure/grpc/functions"
	"go-challenger/infrastructure/grpc/server"
	"go-challenger/infrastructure/logger"
	"sync"
	"time"
)

type config struct {
	configApp *appConfig.AppConfig
	webServer server.Server
	db        *database.MySQLConnection
	repo      coreRepository.TaskRepository
}

func NewConfig() *config {
	return &config{}
}

func (c *config) WithAppConfig() *config {
	var err error
	c.configApp, err = appConfig.LoadConfig()
	if err != nil {
		logger.Fatal(err)
	}
	return c
}

func (c *config) InitLogger() *config {
	logger.NewZapLogger()
	logger.Infof("Log has been successfully configured")
	return c
}

func (c *config) WithDB() *config {
	db, err := database.NewMySQLConnection(c.configApp.MySQL.Host)
	if err != nil {
		logger.Fatal(err)
	}

	c.db = db
	logger.Infof("DB has been successfully configured")
	return c
}

func (c *config) WithRepository() *config {
	c.repo = repository.NewTaskRepository(c.db)
	logger.Infof("Repository has been successfully configured")
	return c
}

func (c *config) WithWebServer() *config {

	intDuration, err := time.ParseDuration(c.configApp.Application.Server.Timeout + "s")
	if err != nil {
		logger.Fatal(err)
	}

	saveUc := usecase.NewSaveUseCase(c.repo)
	deleteUc := usecase.NewDeleteByIdUseCase(c.repo)
	updateUc := usecase.NewUpdateUseCase(c.repo)
	getUc := usecase.NewFindByIdUseCase(c.repo)
	saveAllUc := usecase.NewSaveAllUseCase(c.repo)
	
	c.webServer = server.NewServer(
		functions.NewGrpcFunctions(deleteUc, saveUc,saveAllUc, updateUc, getUc),
		c.configApp.Application.Server.Port, intDuration*time.Second,
	)
	logger.Infof("Web server has been successfully configurated")
	return c
}

func (c *config) Start(ctx context.Context, wg *sync.WaitGroup) {
	logger.Infof("App running on port %s", c.configApp.Application.Server.Port)
	c.webServer.Listen(ctx, wg)

}
