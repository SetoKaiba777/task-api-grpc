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

func TestSave(t *testing.T) {
	logger.NewZapLogger()
	tt := []struct {
		name          string
		input         *input.TaskInput
		mockedOutput  domain.Task
		expected      string
		mockedError   error
		expectedError error
	}{
		{
			name: "saved successfully",
			input: &input.TaskInput{
				Name:   "Fazer altos nadas",
				Status: "Doing",
			},
			mockedOutput: domain.Task{
				Id:     "1",
				Name:   "Fazer altos nadas",
				Status: "Doing",
			},
			mockedError:   nil,
			expectedError: nil,
			expected: "1",
		},
		{
			name: "connection error",
			input: &input.TaskInput{
				Name:   "Fazer altos nadas",
				Status: "Doing",
			},
			mockedOutput:  domain.Task{},
			mockedError:   errGenericConnection,
			expectedError: errGenericConnection,
			expected: "",
		},
		{
			name: "invalid task status",
			input: &input.TaskInput{
				Name:   "Fazer altos nadas",
				Status: "Invalid",
			},
			mockedOutput: domain.Task{
			},
			mockedError:   domain.ErrInvalidStatus,
			expectedError: domain.ErrInvalidStatus,
			expected: "",
		},
	}
	for _, scenario := range tt {
		t.Run(scenario.name, func(t *testing.T) {
			rMock := &mockedRepository{}
			rMock.On("Save", mock.Anything).Return(scenario.mockedOutput, scenario.mockedError)
			uc := NewSaveUseCase(rMock)
			out, err := uc.Execute(context.TODO(), scenario.input)

			assert.Equal(t, scenario.expectedError, err)
			assert.Equal(t, scenario.expected, out)

		})
	}
}
