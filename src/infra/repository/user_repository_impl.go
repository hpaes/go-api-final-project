package repository

import (
	"context"

	"github.com/hpaes/go-api-final-project/src/core/domain"
	"github.com/hpaes/go-api-final-project/src/core/domain/repository"
	"github.com/hpaes/go-api-final-project/src/infra/database"
)

var _ repository.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	db database.DbUsers
}

func NewUserRepository(db database.DbUsers) *UserRepository {
	return &UserRepository{db: db}
}

// GetUserByEmail implements repository.UserRepository.
func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := ur.db.GetByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

// AddUser implements repository.UserRepository.
func (ur *UserRepository) AddUser(ctx context.Context, user *domain.User) (domain.User, error) {
	u, err := ur.db.Save(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

// DeleteUser implements repository.UserRepository.
func (ur *UserRepository) DeleteUser(ctx context.Context, id int) error {
	err := ur.db.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// GetUserById implements repository.UserRepository.
func (ur *UserRepository) GetUserById(ctx context.Context, id int) (domain.User, error) {
	u, err := ur.db.GetById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

// ListUsers implements repository.UserRepository.
func (ur *UserRepository) ListUsers(ctx context.Context, page int) ([]domain.User, error) {
	us, err := ur.db.List(ctx, page)
	if err != nil {
		return []domain.User{}, err
	}
	return us, nil
}

// UpdateUser implements repository.UserRepository.
func (ur *UserRepository) UpdateUser(ctx context.Context, user *domain.User) (domain.User, error) {
	u, err := ur.db.Update(ctx, user)
	if err != nil {
		return domain.User{}, err
	}
	return u, nil
}
