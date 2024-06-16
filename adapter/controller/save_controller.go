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
	SaveController interface{
		Execute(ctx context.Context, task *input.TaskInput) (string, codes.Code ,error)
	}

	saveController struct {
		uc usecase.SaveUseCaseInterface
	}
)

func NewSaveController(uc usecase.SaveUseCaseInterface) SaveController{
	return &saveController{
		uc: uc,
	}
}

func (c *saveController) Execute(ctx context.Context, task *input.TaskInput) (string,codes.Code ,error) {
	logger.Infof("saveController.Execute has been started...")
	out, err := c.uc.Execute(ctx, task)
	if err != nil{
		logger.Errorf("Error in saveUseCase.Execute: %v", err)
		return "", handler.GetStatus(err),err
	}
	logger.Infof("saveController.Execute has been finished...")
	return out,codes.OK ,err
}