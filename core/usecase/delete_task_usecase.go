package usecase

import (
	"context"
	"go-challenger/core/repository"
	"go-challenger/infrastructure/logger"
)

type (
	DeleteByIdUseCaseInterface interface {
		Execute(ctx context.Context, id string) error
	}

	deleteByIdUseCase struct{
		repository repository.TaskRepository
	}
)

func NewDeleteByIdUseCase(repository repository.TaskRepository) DeleteByIdUseCaseInterface{
	return &deleteByIdUseCase{
		repository: repository,
	}
}

func (uc * deleteByIdUseCase) Execute(ctx context.Context,id string) error{
	logger.Infof("deleteByIdUseCase.Execute has been started...")
	_, err:= uc.repository.FindById(ctx,id)
	if err!= nil {
		return err
	}

	if err := uc.repository.Delete(ctx,id);err!=nil{
		logger.Errorf("Error in repository", err)
		return err
	}
	logger.Infof("delete task is successful!")
	logger.Infof("deleteByIdUseCase.Execute has been finished...")
	return nil	
}