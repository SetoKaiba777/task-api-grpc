package usecase

import (
	"context"
	"go-challenger/core/domain"
	"go-challenger/core/repository"
	"go-challenger/infrastructure/logger"
)

type (
	FindByIdUseCaseInterface interface {
		Execute(ctx context.Context, id string) (domain.Task, error)
	}
	
	findByIdUseCase struct {
		repository repository.TaskRepository
	}
)

func NewFindByIdUseCase(repository repository.TaskRepository) FindByIdUseCaseInterface{
	return &findByIdUseCase{
		repository: repository,
	}
}

func (uc * findByIdUseCase) Execute(ctx context.Context, id string) (domain.Task, error){
	logger.Infof("findByIdUseCase.Execute has been started...")
	task, err:= uc.repository.FindById(ctx,id)
	if err!= nil {
		logger.Errorf("Error in repository", err)
		return domain.Task{}, err
	}
	logger.Infof("update task is successful!")
	logger.Infof("findByIdUseCase.Execute has been started...")
	return task, nil
}
