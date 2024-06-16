package usecase

import (
	"context"
	"go-challenger/core/domain"
	"go-challenger/core/repository"
	"go-challenger/core/usecase/input"
	"go-challenger/infrastructure/logger"
)

type(
	SaveUseCaseInterface interface {
	Execute(ctx context.Context, i *input.TaskInput )  (string,error)
}
	saveUseCase struct{
		repository repository.TaskRepository
	}
)

func NewSaveUseCase(repository repository.TaskRepository) SaveUseCaseInterface{
	return &saveUseCase{repository: repository}
}

func (uc *saveUseCase) Execute(ctx context.Context, i *input.TaskInput ) (string,error){
	logger.Infof("saveUseCase.Execute has been started...")
	task, err := domain.NewTask().
						WithId(i.Id).
						WithName(i.Name).
						WithStatus(i.Status).
						Build()
	if err != nil{
		logger.Errorf("Invalid task status: %v", err)
		return "", err
	}

	out, err := uc.repository.Save(ctx,*task)
	if err != nil{
		logger.Errorf("Error in repository: %v", err)
		return "",err
	}
	logger.Infof("Id %s", out)
	logger.Infof("saveUseCase.Execute has been finished...")
	return out.Id ,nil
}