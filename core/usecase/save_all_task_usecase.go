package usecase

import (
	"context"
	"go-challenger/core/domain"
	"go-challenger/core/repository"
	"go-challenger/core/usecase/input"
	"go-challenger/infrastructure/logger"
)

type(
	SaveAllUseCaseInterface interface {
	Execute(ctx context.Context, i []input.TaskInput )  error
}
	saveAllUseCase struct{
		repository repository.TaskRepository
	}
)

func NewSaveAllUseCase(repository repository.TaskRepository) SaveAllUseCaseInterface{
	return &saveAllUseCase{repository: repository}
}

func (uc *saveAllUseCase) Execute(ctx context.Context, i []input.TaskInput ) error{
	logger.Infof("saveAllUseCase.Execute has been started...")
	var taskList []domain.Task
	for _,input := range i {
		task, err := domain.NewTask().
		WithId(input.Id).
		WithName(input.Name).
		WithStatus(input.Status).
		Build()
		
		if err != nil{
			logger.Errorf("Invalid task status: %v", err)
			return err
		}

		taskList=append(taskList, *task)
	} 

	err := uc.repository.SaveAll(ctx,taskList)
	if err != nil{
		logger.Errorf("Error in repository: %v", err)
		return err
	}
	logger.Infof("saveAllUseCase.Execute has been finished...")
	return nil
}