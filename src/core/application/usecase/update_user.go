package usecase

import (
	"context"

	"github.com/hpaes/go-api-final-project/src/core/domain"
	"github.com/hpaes/go-api-final-project/src/core/domain/repository"
	"github.com/hpaes/go-api-final-project/src/infra/logger"
)

type (
	UpdateUser interface {
		Execute(ctx context.Context, input *UpdateUserInput) (*UpdateUserOutput, error)
	}
	updateUser struct {
		userRepository repository.UserRepository
		logger         logger.LoggerService
	}
	UpdateUserInput struct {
		Id    int
		Name  string
		Email string
		Age   int
	}
	UpdateUserOutput struct {
		Id    int
		Name  string
		Email string
		Age   int
	}
)

func NewUpdateUser(userRepository repository.UserRepository, logger logger.LoggerService) UpdateUser {
	return &updateUser{
		userRepository: userRepository,
		logger:         logger,
	}
}

// Execute implements UpdateUser.
func (u *updateUser) Execute(ctx context.Context, input *UpdateUserInput) (*UpdateUserOutput, error) {
	user, err := domain.Create(input.Name, input.Email, input.Age)
	if err != nil {
		u.logger.WithFields(logger.Fields{
			"Name":  input.Name,
			"Email": input.Email,
			"Age":   input.Age,
		}).Error("Failed to create user")
		return nil, err
	}
	userToUpdate, err := u.userRepository.GetUserById(ctx, input.Id)
	if err != nil {
		u.logger.WithFields(logger.Fields{
			"Id": input.Id,
		}).Error("Failed to get user by id")
		return nil, err
	}

	userToUpdate.Age = user.Age
	userToUpdate.Name = user.Name
	userToUpdate.Email = user.Email

	updatedUser, err := u.userRepository.UpdateUser(ctx, &userToUpdate)
	if err != nil {
		u.logger.WithFields(logger.Fields{
			"Id":    userToUpdate.Id,
			"Name":  userToUpdate.Name,
			"Email": userToUpdate.Email,
			"Age":   userToUpdate.Age,
		}).Error("Failed to update user")
		return nil, err
	}

	u.logger.WithFields(logger.Fields{
		"Id":    updatedUser.Id,
		"Name":  updatedUser.Name,
		"Email": updatedUser.Email,
		"Age":   updatedUser.Age,
	}).Info("Updated user successfully")

	return &UpdateUserOutput{
		Id:    updatedUser.Id,
		Name:  updatedUser.Name,
		Email: updatedUser.Email.Value(),
		Age:   updatedUser.Age,
	}, nil
}
