package usecase

import (
	"context"

	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
	"github.com/hpaes/go-api-final-project/src/core/domain/repository"
	"github.com/hpaes/go-api-final-project/src/infra/logger"
)

type (
	RemoveUser interface {
		Execute(ctx context.Context, input *RemoveUserInput) error
	}
	removeUserUsecase struct {
		userRepository repository.UserRepository
		logger         logger.LoggerService
	}

	RemoveUserInput struct {
		Id int
	}
)

func NewRemoveUser(userRepository repository.UserRepository, logger logger.LoggerService) RemoveUser {
	return &removeUserUsecase{
		userRepository: userRepository,
		logger:         logger,
	}
}

// Execute implements RemoveUser.
func (r *removeUserUsecase) Execute(ctx context.Context, input *RemoveUserInput) error {
	if input.Id == 0 {
		r.logger.WithFields(logger.Fields{
			"Id": input.Id,
		}).Error("User Id is required")
		return custom_errors.NewUserIdIsRequiredError()
	}

	err := r.userRepository.DeleteUser(ctx, input.Id)
	if err != nil {
		r.logger.WithFields(logger.Fields{
			"Id": input.Id,
		}).Error("Failed to remove user")
		return err
	}

	r.logger.WithFields(logger.Fields{
		"Id": input.Id,
	}).Info("Removed user successfully")

	return nil
}
