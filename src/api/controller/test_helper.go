package controller

import (
	"context"

	"github.com/hpaes/go-api-final-project/src/core/application/usecase"
	"github.com/hpaes/go-api-final-project/src/core/domain"
	"github.com/stretchr/testify/mock"
)

type (
	mockUsecaseFactory struct {
		mock.Mock
	}
	mockRegisterUser struct {
		mock.Mock
	}
	mockGetUserDetail struct {
		mock.Mock
	}
	mockRemoveUser struct {
		mock.Mock
	}
	mockListUsers struct {
		mock.Mock
	}
	mockUpdateUser struct {
		mock.Mock
	}
	mockUserRepository struct {
		mock.Mock
	}
)

func (m *mockRegisterUser) Execute(ctx context.Context, input *usecase.RegisterUserInput) (*usecase.RegisterUserOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*usecase.RegisterUserOutput), args.Error(1)
}

func (m *mockGetUserDetail) Execute(ctx context.Context, input *usecase.GetUserDetailInput) (domain.User, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *mockRemoveUser) Execute(ctx context.Context, input *usecase.RemoveUserInput) error {
	return m.Called(ctx, input).Error(0)
}

func (m *mockListUsers) Execute(ctx context.Context, input *usecase.ListUsersInput) ([]domain.User, error) {
	args := m.Called(ctx, input)
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *mockUpdateUser) Execute(ctx context.Context, input *usecase.UpdateUserInput) (*usecase.UpdateUserOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*usecase.UpdateUserOutput), args.Error(1)
}
