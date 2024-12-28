package usecase

import (
	"context"
	"testing"

	"github.com/hpaes/go-api-final-project/src/core/domain"
	"github.com/hpaes/go-api-final-project/src/core/domain/vo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListUsers(t *testing.T) {
	email, _ := vo.NewEmail("email@test.com")
	dbReturn := []domain.User{
		{
			Id:    1,
			Name:  "Test User 1",
			Email: *email,
			Age:   10,
		},
	}

	tt := []struct {
		name     string
		input    *ListUsersInput
		expected []domain.User
		err      error
	}{
		{
			name:     "Success",
			input:    &ListUsersInput{Page: 1},
			expected: dbReturn,
			err:      nil,
		},
		{
			name:     "Page is 0",
			input:    &ListUsersInput{Page: 0},
			expected: nil,
			err:      databaseErr,
		},
		{
			name:     "Page is negative",
			input:    &ListUsersInput{Page: -1},
			expected: nil,
			err:      databaseErr,
		},
		{
			name:     "not found",
			input:    &ListUsersInput{Page: 1},
			expected: []domain.User{},
			err:      nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			loggerMock := NewMockLogger()
			mockDB := NewMockDatabase()
			usecase := NewListUsers(mockDB, loggerMock)

			mockDB.On("ListUsers", mock.Anything, tc.input.Page).Return(tc.expected, tc.err)
			loggerMock.On("WithFields", mock.Anything).Return(loggerMock).Once()
			loggerMock.On("Info", mock.Anything).Return().Maybe()
			loggerMock.On("Error", mock.Anything).Return().Maybe()

			ctx := context.TODO()
			users, err := usecase.Execute(ctx, tc.input)

			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.expected, users)
			mockDB.AssertExpectations(t)
		})
	}
}
