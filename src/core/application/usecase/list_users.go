package usecase

import (
	"context"

	"github.com/hpaes/go-api-final-project/src/core/domain"
	"github.com/hpaes/go-api-final-project/src/core/domain/repository"
	"github.com/hpaes/go-api-final-project/src/infra/logger"
)

type (
	ListUsersInput struct {
		Page int
	}

	listUsers struct {
		userRepository repository.UserRepository
		logger         logger.LoggerService
	}

	ListUsers interface {
		Execute(ctx context.Context, input *ListUsersInput) ([]domain.User, error)
	}
)

func NewListUsers(userRepository repository.UserRepository, logger logger.LoggerService) ListUsers {
	return &listUsers{
		userRepository: userRepository,
		logger:         logger,
	}
}

// Execute implements ListUsers.
func (l *listUsers) Execute(ctx context.Context, input *ListUsersInput) ([]domain.User, error) {
	users, err := l.userRepository.ListUsers(ctx, input.Page)
	if err != nil {
		l.logger.WithFields(logger.Fields{
			"Page": input.Page,
		}).Error("Failed to list users")
		return nil, err
	}

	if len(users) == 0 {
		return []domain.User{}, nil
	}

	l.logger.WithFields(logger.Fields{
		"Page": input.Page,
	}).Info("List users success")
	return users, nil
}
