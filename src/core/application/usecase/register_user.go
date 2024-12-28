package usecase

import (
	"context"

	"github.com/hpaes/go-api-final-project/src/core/domain"
	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
	"github.com/hpaes/go-api-final-project/src/core/domain/repository"
	"github.com/hpaes/go-api-final-project/src/infra/logger"
)

type (
	RegisterUser interface {
		Execute(ctx context.Context, input *RegisterUserInput) (*RegisterUserOutput, error)
	}
	registerUser struct {
		userRepository repository.UserRepository
		logger         logger.LoggerService
	}
	RegisterUserInput struct {
		Id    string
		Name  string
		Email string
		Age   int
	}

	RegisterUserOutput struct {
		Id    int
		Name  string
		Email string
		Age   int
	}
)

func NewRegisterUser(userRepository repository.UserRepository, logger logger.LoggerService) RegisterUser {
	return &registerUser{
		userRepository: userRepository,
		logger:         logger,
	}
}

// Execute implements RegisterUser.
func (r *registerUser) Execute(ctx context.Context, input *RegisterUserInput) (*RegisterUserOutput, error) {
	r.logger.Info("Executing RegisterUser use case")
	u, err := domain.Create(input.Name, input.Email, input.Age)
	if err != nil {
		r.logger.WithFields(logger.Fields{
			"Name":  input.Name,
			"Email": input.Email,
			"Age":   input.Age,
		}).Error("Failed to create user")
		return nil, err
	}
	userExists, _ := r.userRepository.GetUserByEmail(ctx, u.Email.Value())
	if userExists.Email.Value() != "" && userExists.Email.Value() == u.Email.Value() {
		r.logger.WithFields(logger.Fields{
			"Email": input.Email,
		}).Error("User already exists")
		return nil, custom_errors.NewUserAlreadyExistsError(u.Email.Value())
	}
	user, err := r.userRepository.AddUser(ctx, u)
	if err != nil {
		r.logger.WithFields(logger.Fields{
			"Name":  input.Name,
			"Email": input.Email,
			"Age":   input.Age,
		}).Error("Failed to register user")
		return nil, err
	}

	r.logger.WithFields(logger.Fields{
		"Id":    user.Id,
		"Name":  user.Name,
		"Email": user.Email.Value(),
		"Age":   user.Age,
	}).Info("User registered successfully")

	return &RegisterUserOutput{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email.Value(),
		Age:   user.Age,
	}, nil

}
