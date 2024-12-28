package domain

import (
	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
	"github.com/hpaes/go-api-final-project/src/core/domain/vo"
)

type User struct {
	Id    int      `json:"id"`
	Name  string   `json:"name"`
	Email vo.Email `json:"email"`
	Age   int      `json:"age"`
}

func newUser(id int, name, email string, age int) (*User, error) {
	if name == "" {
		return nil, custom_errors.NewParameterRequiredError("Name")
	}
	if age <= 0 {
		return nil, custom_errors.NewInvalidParameterError("Age", "Age must be higher than 0")
	}
	emailVo, err := vo.NewEmail(email)
	if err != nil {
		return nil, err
	}
	return &User{
		Id:    id,
		Name:  name,
		Email: *emailVo,
		Age:   age,
	}, nil
}

func Create(name, email string, age int) (*User, error) {
	user, err := newUser(0, name, email, age)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func Restore(id int, name, email string, age int) (*User, error) {
	user, err := newUser(id, name, email, age)
	if err != nil {
		return nil, err
	}
	return user, nil
}
