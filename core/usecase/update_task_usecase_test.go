package usecase

import (
	"context"
	"go-challenger/core/domain"
	"go-challenger/core/usecase/input"
	"go-challenger/infrastructure/logger"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateUseCasae(t *testing.T) {
	logger.NewZapLogger()
	tt := []struct {
		name             string
		input            input.TaskInput
		output           input.TaskInput
		mockUpdateOutput domain.Task
		mockGetOutput    domain.Task
		mockUpdateError  error
		mockGetError     error
		expectedError    error
	}{
		{
			name: "update successfully",
			input: input.TaskInput{
				Id:     "1",
				Name:   "Fazer altos nadas",
				Status: "Doing",
			},
			mockGetOutput: domain.Task{
				Id:     "1",
				Name:   "Fazer altos nadas em casa",
				Status: "Doing",
			},
			mockUpdateOutput: domain.Task{
				Id:     "1",
				Name:   "Fazer altos nadas",
				Status: "Doing",
			},
			mockGetError:    nil,
			mockUpdateError: nil,
			expectedError:   nil,
		},
		{
			name: "connection error",
			input: input.TaskInput{
				Id:     "1",
				Name:   "Fazer altos nadas",
				Status: "Doing",
			},
			mockUpdateOutput: domain.Task{},
			mockGetOutput: domain.Task{},
			mockGetError:     errGenericConnection,
			mockUpdateError:  nil,
			expectedError:    errGenericConnection,
		},
	}
	for _, scenario := range tt {
		t.Run(scenario.name, func(t *testing.T) {
			rMock := &mockedRepository{}
			rMock.On("FindById", mock.Anything).Return(scenario.mockGetOutput, scenario.mockGetError)
			rMock.On("Update", mock.Anything).Return(scenario.mockUpdateOutput, scenario.mockUpdateError)
			uc := NewUpdateUseCase(rMock)
			i := scenario.input
			_, err := uc.Execute(context.TODO(), &i)
			assert.Equal(t, scenario.expectedError, err)
		})
	}
}
