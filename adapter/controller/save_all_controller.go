package controller

import (
	"context"
	"go-challenger/adapter/handler"
	"go-challenger/core/usecase"
	"go-challenger/core/usecase/input"
	"go-challenger/infrastructure/logger"

	"google.golang.org/grpc/codes"
)

type (
	SaveAllController interface{
		Execute(ctx context.Context, task *[]input.TaskInput) (codes.Code ,error)
	}

	saveAllController struct {
		uc usecase.SaveAllUseCaseInterface
	}
)

func NewSaveAllController(uc usecase.SaveAllUseCaseInterface) SaveAllController{
	return &saveAllController{
		uc: uc,
	}
}

func (c *saveAllController) Execute(ctx context.Context, tasks *[]input.TaskInput) (codes.Code ,error) {
	logger.Infof("saveAllController.Execute has been started...")
	err := c.uc.Execute(ctx, *tasks)
	if err != nil{
		logger.Errorf("Error in saveUseCase.Execute: %v", err)
		return handler.GetStatus(err),err
	}
	logger.Infof("saveController.Execute has been finished...")
	return codes.OK ,err
}