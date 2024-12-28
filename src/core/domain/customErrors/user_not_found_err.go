package custom_errors

import "fmt"

type UserNotFoundError struct {
	id    string
	param string
}

func NewUserNotFoundError(id, param string) *UserNotFoundError {
	return &UserNotFoundError{id: id, param: param}
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("%s not found, id: %s", e.param, e.id)
}
