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

func TestRegisterUser(t *testing.T) {
	email, _ := vo.NewEmail("email@test.com")
	domainExpected := domain.User{
		Id:    1,
		Name:  "Test User",
		Email: *email,
		Age:   10,
	}

	registerInput := &RegisterUserInput{
		Name:  "Test User",
		Email: "email@test.com",
		Age:   10,
	}

	registerExpected := &RegisterUserOutput{
		Id:    1,
		Name:  "Test User",
		Email: "email@test.com",
		Age:   10,
	}

	tt := []struct {
		name               string
		input              *RegisterUserInput
		expected           *RegisterUserOutput
		addDbExpected      domain.User
		getByEmailExpected domain.User
		err                error
	}{
		{
			name:               "Success",
			input:              registerInput,
			expected:           registerExpected,
			addDbExpected:      domainExpected,
			getByEmailExpected: domain.User{},
			err:                nil,
		},
		{
			name:               "User already exists",
			input:              registerInput,
			expected:           nil,
			getByEmailExpected: domainExpected,
			addDbExpected:      domain.User{},
			err:                custom_errors.NewUserAlreadyExistsError(registerInput.Email),
		},
		{
			name:               "Error",
			input:              registerInput,
			expected:           nil,
			addDbExpected:      domain.User{},
			getByEmailExpected: domain.User{},
			err:                databaseErr,
		},
		{
			name:               "error on create user",
			input:              &RegisterUserInput{},
			expected:           nil,
			addDbExpected:      domain.User{},
			getByEmailExpected: domain.User{},
			err:                custom_errors.NewParameterRequiredError("Name"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			loggerMock := NewMockLogger()
			mockDB := NewMockDatabase()
			usecase := NewRegisterUser(mockDB, loggerMock)

			if tc.input.Name != "" {
				mockDB.On("GetUserByEmail", mock.Anything, tc.input.Email).Return(tc.getByEmailExpected, tc.err)
			}

			if tc.getByEmailExpected.Name == "" && tc.input.Name != "" {
				mockDB.On("AddUser", mock.Anything, mock.Anything).Return(tc.addDbExpected, tc.err)
			}

			loggerMock.On("WithFields", mock.Anything).Return(loggerMock).Once()
			loggerMock.On("Info", mock.Anything).Return().Maybe()
			loggerMock.On("Error", mock.Anything).Return().Maybe()

			ctx := context.TODO()
			user, err := usecase.Execute(ctx, tc.input)

			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.expected, user)
			mockDB.AssertExpectations(t)
		})
	}

}
