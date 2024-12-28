package domain

import (
	"testing"

	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// Test valid user creation
	user, err := Create("John Doe", "john.doe@example.com", 30)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john.doe@example.com", user.Email.Value())
	assert.Equal(t, 30, user.Age)

	// Test user creation with empty name
	user, err = Create("", "john.doe@example.com", 30)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.IsType(t, &custom_errors.ParameterRequiredError{}, err)

	// Test user creation with invalid age
	user, err = Create("John Doe", "john.doe@example.com", -1)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.IsType(t, &custom_errors.InvalidParameterError{}, err)

	// Test user creation with invalid email
	user, err = Create("John Doe", "invalid-email", 30)
	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestRestoreUser(t *testing.T) {
	// Test valid user restoration
	user, err := Restore(1, "John Doe", "john.doe@example.com", 30)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, 1, user.Id)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john.doe@example.com", user.Email.Value())
	assert.Equal(t, 30, user.Age)

	// Test user restoration with empty name
	user, err = Restore(1, "", "john.doe@example.com", 30)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.IsType(t, &custom_errors.ParameterRequiredError{}, err)

	// Test user restoration with invalid age
	user, err = Restore(1, "John Doe", "john.doe@example.com", -1)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.IsType(t, &custom_errors.InvalidParameterError{}, err)

	// Test user restoration with invalid email
	user, err = Restore(1, "John Doe", "invalid-email", 30)
	assert.Error(t, err)
	assert.Nil(t, user)
}
