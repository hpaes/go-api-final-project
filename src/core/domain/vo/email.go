package vo

import (
	"encoding/json"
	"regexp"

	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
)

type Email struct {
	value string
}

const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

// NewEmail creates a new Email instance and validates the format.
func NewEmail(value string) (*Email, error) {
	if !isValidEmail(value) {
		return nil, custom_errors.NewInvalidParameterError("Email", "Invalid email format")
	}
	return &Email{value: value}, nil
}

// Value returns the email value as a string.
func (e Email) Value() string {
	return e.value
}

// MarshalJSON marshals the email value to JSON.
func (e Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.value)
}

func isValidEmail(value string) bool {
	regex, err := regexp.Compile(emailRegexPattern)
	if err != nil {
		return false
	}
	return regex.MatchString(value)
}

// Validate checks if the email format is valid.
func (e Email) Validate() error {
	if !isValidEmail(e.value) {
		return custom_errors.NewInvalidParameterError("Email", "Invalid email format")
	}
	return nil
}
