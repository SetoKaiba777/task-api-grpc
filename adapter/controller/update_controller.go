package controller

import (
	"context"
	"go-challenger/adapter/handler"
	"go-challenger/core/domain"
	"go-challenger/core/usecase"
	"go-challenger/core/usecase/input"
	"go-challenger/infrastructure/logger"

	"google.golang.org/grpc/codes"
)

type (
	UpdateController interface {
		Execute(ctx context.Context, task *input.TaskInput) (domain.Task,codes.Code ,error)
	}

	updateController struct {
		uc usecase.UpdateUseCaseInterface
	}
)

func NewSaveUpdateController(uc usecase.UpdateUseCaseInterface) UpdateController {
	return &updateController{
		uc: uc,
	}
}

func (c *updateController) Execute(ctx context.Context, task *input.TaskInput) (domain.Task, codes.Code ,error) {
	logger.Infof("updateController.Execute has been started...")
	out, err := c.uc.Execute(ctx, task)
	if err != nil {
		logger.Errorf("Error in saveUseCase.Execute: %v", err)
		return domain.Task{},handler.GetStatus(err) ,err
	}
	logger.Infof("updateController.Execute has been finished...")
	return out, codes.OK ,err
}