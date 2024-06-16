package usecase

import (
	"context"
	"go-challenger/core/domain"
	"go-challenger/core/repository"
	"go-challenger/core/usecase/input"
	"go-challenger/infrastructure/logger"
)

type(
	UpdateUseCaseInterface interface {
	Execute(ctx context.Context, i *input.TaskInput)  (domain.Task,error)
}
	updateUseCase struct{
		repository repository.TaskRepository
	}
)

func NewUpdateUseCase(repository repository.TaskRepository) UpdateUseCaseInterface{
	return &updateUseCase{repository: repository}
}

func (uc *updateUseCase) Execute(ctx context.Context, i *input.TaskInput) (domain.Task, error){
	logger.Infof("updateUseCase.Execute has been started...")
	logger.Infof("find by id %s in repository", i.Id)
	out, err := uc.repository.FindById(ctx,i.Id)
	if err != nil{
		switch(err){
		case domain.ErrNotFoundTask:
			logger.Errorf("not found id %s in repository", i.Id)
			return domain.Task{},err
		default:
			logger.Errorf("repository error: %v",err)
			return	domain.Task{}, err
		}
	}
	if i.Status != "" {
		task, err := domain.NewTask().
		WithId(i.Id).
		WithName(i.Name).
		WithStatus(i.Status).
		Build()
		if err != nil{
			logger.Errorf("invalid id status: %s", i.Status)
			return domain.Task{}, err
		}
		
		out, err := uc.repository.Update(ctx,*task)
		if err != nil{
			logger.Errorf("Error in repository", err)
			return domain.Task{},err
		}
		logger.Infof("update task is successful!")
		logger.Infof("updateUseCase.Execute has been finished...")
		return out, nil
	}

	task := domain.NewTask().
	WithId(i.Id).
	WithName(i.Name).
	WithStatus(out.Status)

	out, err = uc.repository.Update(ctx,*task)
	if err != nil{
		return domain.Task{},err
	}
	logger.Infof("updateUseCase.Execute has been finished...")
	return out,nil
}