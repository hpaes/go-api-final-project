package usecase

import (
	"context"
	"testing"

	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRemoveUser(t *testing.T) {
	input := &RemoveUserInput{Id: 1}

	tc := []struct {
		name  string
		input *RemoveUserInput
		err   error
	}{
		{
			name:  "success",
			input: input,
			err:   nil,
		},
		{
			name:  "id not informed",
			input: &RemoveUserInput{Id: 0},
			err:   custom_errors.NewUserIdIsRequiredError(),
		},
		{
			name:  "error on delete user",
			input: input,
			err:   databaseErr,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			loggerMock := NewMockLogger()
			mockDB := NewMockDatabase()
			usecase := NewRemoveUser(mockDB, loggerMock)

			if tt.input.Id != 0 {
				mockDB.On("DeleteUser", mock.Anything, tt.input.Id).Return(tt.err)
			}
			loggerMock.On("WithFields", mock.Anything).Return(loggerMock).Once()
			loggerMock.On("Info", mock.Anything).Return().Maybe()
			loggerMock.On("Error", mock.Anything).Return().Maybe()

			ctx := context.Background()
			err := usecase.Execute(ctx, tt.input)

			assert.Equal(t, tt.err, err)
			mockDB.AssertExpectations(t)
		})
	}
}
