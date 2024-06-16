package controller

import (
	"context"
	"go-challenger/adapter/handler"
	"go-challenger/core/domain"
	"go-challenger/core/usecase"
	"go-challenger/infrastructure/logger"

	"google.golang.org/grpc/codes"
)

type (
	GetController interface {
		Execute(ctx context.Context, id string) (*domain.Task, codes.Code, error)
	}
	getController struct {
		uc usecase.FindByIdUseCaseInterface
	}
)

func NewGetController(uc usecase.FindByIdUseCaseInterface) GetController{
	return &getController{
		uc: uc,
	}
}

func (c *getController) Execute(ctx context.Context, id string) (*domain.Task,codes.Code, error){
	logger.Infof("getController.Execute has been started...")
	task, err := c.uc.Execute(ctx, id)
	if err != nil {
		logger.Errorf("Error in findByIdUseCase.Execute: %v", err)
		return nil, handler.GetStatus(err) ,err
	}
	logger.Infof("getController.Execute has been finished...")
	return &task, codes.OK, nil
}