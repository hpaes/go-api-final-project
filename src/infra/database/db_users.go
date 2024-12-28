package database

import (
	"context"

	"github.com/hpaes/go-api-final-project/src/core/domain"
)

type DbUsers interface {
	Save(ctx context.Context, user *domain.User) (domain.User, error)
	GetById(ctx context.Context, id int) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	Update(ctx context.Context, user *domain.User) (domain.User, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, page int) ([]domain.User, error)
}
