package vo

import (
	"regexp"

	"encoding/json"

	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	if !isValidEmail(value) {
		return nil, custom_errors.NewInvalidParameterError("Email", "Invalid email format")
	}
	return &Email{value: value}, nil
}

func (e Email) Value() string {
	return e.value
}

func (e Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.value)
}

func isValidEmail(value string) bool {
	regex, _ := regexp.Compile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(value)
}
