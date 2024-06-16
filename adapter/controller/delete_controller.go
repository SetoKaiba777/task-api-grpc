package controller

import (
	"context"
	"go-challenger/adapter/handler"
	"go-challenger/core/usecase"
	"go-challenger/infrastructure/logger"

	"google.golang.org/grpc/codes"
)

type (
	DeleteController interface {
		Execute(ctx context.Context, id string) (codes.Code,error)
	}
	deleteController struct {
		uc usecase.DeleteByIdUseCaseInterface
	}
)

func NewDeleteController(uc usecase.DeleteByIdUseCaseInterface) DeleteController {
	return &deleteController{
		uc: uc,
	}
}

func (c *deleteController) Execute(ctx context.Context, id string) (codes.Code, error){
	logger.Infof("deleteController.Execute has been started...")
	if err:=c.uc.Execute(ctx, id); err != nil{
		logger.Errorf("Error in deleteByIdUseCase.Execute: %v", err)
		return handler.GetStatus(err), err
	}
	logger.Infof("deleteController.Execute has been finished...")
	return codes.OK, nil
}