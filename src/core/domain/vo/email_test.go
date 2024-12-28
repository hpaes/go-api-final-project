package vo

import (
	"testing"

	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
	"github.com/stretchr/testify/assert"
)

func TestNewEmail(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"test@example.com", "test@example.com", false},
		{"invalid-email", "", true},
		{"another.test@domain.co", "another.test@domain.co", false},
		{"@missingusername.com", "", true},
	}

	for _, test := range tests {
		email, err := NewEmail(test.input)
		if test.hasError {
			assert.Error(t, err)
			assert.Nil(t, email)
			assert.IsType(t, &custom_errors.InvalidParameterError{}, err)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, email)
			assert.Equal(t, test.expected, email.Value())
		}
	}
}

func TestEmail_MarshalJSON(t *testing.T) {
	email, err := NewEmail("test@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, email)

	jsonData, err := email.MarshalJSON()
	assert.NoError(t, err)
	assert.JSONEq(t, `"test@example.com"`, string(jsonData))
}
