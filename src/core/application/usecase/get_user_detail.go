package usecase

import (
	"context"

	"github.com/hpaes/go-api-final-project/src/core/domain"
	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
	"github.com/hpaes/go-api-final-project/src/core/domain/repository"
	"github.com/hpaes/go-api-final-project/src/infra/logger"
)

type (
	GetUserDetailInput struct {
		Id int
	}

	GetUserDetailOutput struct {
		Id    int
		Name  string
		Email string
		Age   int
	}
	getUserDetail struct {
		userRepository repository.UserRepository
		logger         logger.LoggerService
	}

	GetUserDetail interface {
		Execute(ctx context.Context, input *GetUserDetailInput) (domain.User, error)
	}
)

func NewGetUserDetail(userRepository repository.UserRepository, logger logger.LoggerService) GetUserDetail {
	return &getUserDetail{
		userRepository: userRepository,
		logger:         logger,
	}
}

// Execute implements GetUserDetail.
func (g *getUserDetail) Execute(ctx context.Context, input *GetUserDetailInput) (domain.User, error) {
	if input.Id == 0 {
		g.logger.WithFields(logger.Fields{
			"Id": input.Id,
		}).Error("User Id is required")
		return domain.User{}, custom_errors.NewUserIdIsRequiredError()
	}
	user, err := g.userRepository.GetUserById(ctx, input.Id)
	if err != nil {
		g.logger.WithFields(logger.Fields{
			"Id": input.Id,
		}).Error("Failed to get user detail")
		return domain.User{}, err
	}

	g.logger.WithFields(logger.Fields{
		"Id":    user.Id,
		"Name":  user.Name,
		"Email": user.Email,
		"Age":   user.Age,
	}).Info("Get user detail success")
	return user, nil
}
