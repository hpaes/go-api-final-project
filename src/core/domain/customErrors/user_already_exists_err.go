package custom_errors

import "fmt"

type UserAlreadyExistsError struct {
	email string
}

func NewUserAlreadyExistsError(email string) *UserAlreadyExistsError {
	return &UserAlreadyExistsError{email: email}
}

func (e *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("User with email  %s  already exists", e.email)
}
