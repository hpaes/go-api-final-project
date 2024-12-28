package usecase

import (
	"context"
	"fmt"

	"github.com/hpaes/go-api-final-project/src/core/domain"
	"github.com/hpaes/go-api-final-project/src/infra/logger"
	"github.com/stretchr/testify/mock"
)

type databaseMock struct {
	mock.Mock
}

var databaseErr = fmt.Errorf("database error")

func NewMockDatabase() *databaseMock {
	return &databaseMock{}
}

func (m *databaseMock) AddUser(ctx context.Context, p *domain.User) (domain.User, error) {
	args := m.Called(ctx, p)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *databaseMock) GetUserById(ctx context.Context, id int) (domain.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *databaseMock) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *databaseMock) UpdateUser(ctx context.Context, p *domain.User) (domain.User, error) {
	args := m.Called(ctx, p)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *databaseMock) DeleteUser(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *databaseMock) ListUsers(ctx context.Context, page int) ([]domain.User, error) {
	args := m.Called(ctx, page)
	return args.Get(0).([]domain.User), args.Error(1)
}

type loggerMock struct {
	mock.Mock
}

func NewMockLogger() *loggerMock {
	return &loggerMock{}
}

func (m *loggerMock) Info(format string, args ...interface{}) {
	m.Called(args...)
}

func (m *loggerMock) Warn(format string, args ...interface{}) {
	m.Called(args...)
}

func (m *loggerMock) Error(format string, args ...interface{}) {
	m.Called(args...)
}

func (m *loggerMock) Fatal(format string, args ...interface{}) {
	m.Called(args...)
}

func (m *loggerMock) WithFields(fields logger.Fields) logger.LoggerService {
	args := m.Called(fields)
	return args.Get(0).(logger.LoggerService)
}
