package usecase

import (
	"context"
	"testing"

	"github.com/hpaes/go-api-final-project/src/core/domain"
	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
	"github.com/hpaes/go-api-final-project/src/core/domain/vo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateUser(t *testing.T) {
	input := &UpdateUserInput{
		Id:    1,
		Name:  "John Doe",
		Email: "johnDoe1@email.com",
		Age:   30,
	}

	email, _ := vo.NewEmail("johnDoe@email.com")
	expectedGetUserById := domain.User{
		Id:    1,
		Name:  "John Doe",
		Email: *email,
		Age:   29,
	}
	email1, _ := vo.NewEmail("johnDoe1@email.com")
	expectedUpdateUserDomain := domain.User{
		Id:    1,
		Name:  "John Doe",
		Email: *email1,
		Age:   30,
	}
	expectedUpdateUserOutput := &UpdateUserOutput{
		Id:    1,
		Name:  "John Doe",
		Email: "johnDoe1@email.com",
		Age:   30,
	}

	tests := []struct {
		name                     string
		input                    *UpdateUserInput
		expectedGetUserById      domain.User
		expectedUpdateUserDomain domain.User
		expectedUpdateUserOutput *UpdateUserOutput
		err                      error
		updateErr                error
		getUserByIdErr           error
		mockGetUserById          bool
		mockUpdateUser           bool
	}{
		{
			name:                     "success",
			input:                    input,
			expectedGetUserById:      expectedGetUserById,
			expectedUpdateUserDomain: expectedUpdateUserDomain,
			expectedUpdateUserOutput: expectedUpdateUserOutput,
			err:                      nil,
			updateErr:                nil,
			getUserByIdErr:           nil,
			mockGetUserById:          true,
			mockUpdateUser:           true,
		},
		{
			name:                     "error on create user",
			input:                    &UpdateUserInput{Name: "John Doe", Age: 30},
			expectedGetUserById:      domain.User{},
			expectedUpdateUserDomain: domain.User{},
			expectedUpdateUserOutput: nil,
			err:                      custom_errors.NewInvalidParameterError("Email", "Invalid email format"),
			updateErr:                nil,
			getUserByIdErr:           nil,
			mockGetUserById:          false,
			mockUpdateUser:           false,
		},
		{
			name:                     "error on get user",
			input:                    input,
			expectedGetUserById:      domain.User{},
			expectedUpdateUserDomain: domain.User{},
			expectedUpdateUserOutput: nil,
			err:                      databaseErr,
			updateErr:                nil,
			getUserByIdErr:           databaseErr,
			mockGetUserById:          true,
			mockUpdateUser:           false,
		},
		{
			name:                     "error on user not found",
			input:                    input,
			expectedGetUserById:      domain.User{},
			expectedUpdateUserDomain: domain.User{},
			expectedUpdateUserOutput: nil,
			err:                      custom_errors.NewUserNotFoundError("User", "1"),
			updateErr:                nil,
			getUserByIdErr:           custom_errors.NewUserNotFoundError("User", "1"),
			mockGetUserById:          true,
			mockUpdateUser:           false,
		},
		{
			name:                     "error on update user",
			input:                    input,
			expectedGetUserById:      expectedGetUserById,
			expectedUpdateUserDomain: domain.User{},
			expectedUpdateUserOutput: nil,
			err:                      databaseErr,
			updateErr:                databaseErr,
			getUserByIdErr:           nil,
			mockGetUserById:          true,
			mockUpdateUser:           true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dbMock := NewMockDatabase()
			loggerMock := NewMockLogger()
			usecase := NewUpdateUser(dbMock, loggerMock)

			if tc.mockGetUserById {
				dbMock.On("GetUserById", mock.Anything, input.Id).Return(tc.expectedGetUserById, tc.getUserByIdErr)
			}

			if tc.mockUpdateUser {
				dbMock.On("UpdateUser", mock.Anything, mock.Anything).Return(tc.expectedUpdateUserDomain, tc.updateErr)
			}

			loggerMock.On("WithFields", mock.Anything).Return(loggerMock).Once()
			loggerMock.On("Info", mock.Anything).Return().Maybe()
			loggerMock.On("Error", mock.Anything).Return().Maybe()

			ctx := context.Background()
			user, err := usecase.Execute(ctx, tc.input)

			assert.Equal(t, tc.expectedUpdateUserOutput, user)
			assert.Equal(t, tc.err, err)
			dbMock.AssertExpectations(t)
			loggerMock.AssertExpectations(t)
		})
	}
}
