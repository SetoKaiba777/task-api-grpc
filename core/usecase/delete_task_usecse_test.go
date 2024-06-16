package usecase

import (
	"context"
	"errors"
	"go-challenger/core/domain"
	"go-challenger/infrastructure/logger"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	errGenericConnection = errors.New("genereic connection error")
)

func TestDeleteUseCase(t *testing.T) {
	logger.NewZapLogger()
	tt := []tableTest{
		{
			name: "deleted user",
			input: "1",
			output: domain.Task{
				Id: "1",
				Name: "Fazer altos nadas",
				Status: "Doing",
			},
			mockedError: nil,
			expectedError: nil,
		},
		{
			name: "connection error",
			input: "1",
			output: domain.Task{},
			mockedError: errGenericConnection,
			expectedError: errGenericConnection,
		},
		{
			name: "not found user to delete",
			input:"1",
			output: domain.Task{},
			mockedError:domain.ErrNotFoundTask,
			expectedError: domain.ErrNotFoundTask,
		},
	}
	for _, scenario := range tt{
		t.Run(scenario.name, func(t * testing.T){
			rMock := &mockedRepository{}
			rMock.On("FindById",mock.Anything).Return(scenario.output.(domain.Task),scenario.mockedError)
			rMock.On("Delete",mock.Anything).Return(nil)

			uc := NewDeleteByIdUseCase(rMock)
			err := uc.Execute(context.TODO(), scenario.input.(string))

			assert.Equal(t,scenario.expectedError,err)
		})
	}
}