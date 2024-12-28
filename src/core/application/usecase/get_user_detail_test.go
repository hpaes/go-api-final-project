package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/hpaes/go-api-final-project/src/core/domain"
	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
	"github.com/hpaes/go-api-final-project/src/core/domain/vo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserDetail(t *testing.T) {
	email, _ := vo.NewEmail("johnDoe@email.com")
	input := &GetUserDetailInput{Id: 1}
	expected := domain.User{
		Id:    1,
		Name:  "John Doe",
		Email: *email,
		Age:   30,
	}
	databaseErr := errors.New("database error")

	tt := []struct {
		name     string
		input    *GetUserDetailInput
		expected domain.User
		err      error
	}{
		{
			name:     "success",
			input:    input,
			expected: expected,
			err:      nil,
		},
		{
			name:     "id not informed",
			input:    &GetUserDetailInput{Id: 0},
			expected: domain.User{},
			err:      custom_errors.NewUserIdIsRequiredError(),
		},
		{
			name:     "error on get user",
			input:    input,
			expected: domain.User{},
			err:      databaseErr,
		},
		{
			name:     "user not found",
			input:    &GetUserDetailInput{Id: 1},
			expected: domain.User{},
			err:      custom_errors.NewUserNotFoundError("User", "1"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockLogger := NewMockLogger()
			mockDB := NewMockDatabase()
			usecase := NewGetUserDetail(mockDB, mockLogger)

			if tc.input.Id != 0 {
				mockDB.On("GetUserById", mock.Anything, tc.input.Id).Return(tc.expected, tc.err)
			}
			mockLogger.On("WithFields", mock.Anything).Return(mockLogger).Once()
			mockLogger.On("Info", mock.Anything).Return().Maybe()
			mockLogger.On("Error", mock.Anything).Return().Maybe()

			ctx := context.Background()
			user, err := usecase.Execute(ctx, tc.input)

			assert.Equal(t, tc.expected, user)
			assert.Equal(t, tc.err, err)
			mockDB.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
		})
	}
}
