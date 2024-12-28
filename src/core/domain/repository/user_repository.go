package repository

import (
	"context"

	"github.com/hpaes/go-api-final-project/src/core/domain"
)

type UserRepository interface {
	AddUser(ctx context.Context, user *domain.User) (domain.User, error)
	GetUserById(ctx context.Context, id int) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
	ListUsers(ctx context.Context, page int) ([]domain.User, error)
}
