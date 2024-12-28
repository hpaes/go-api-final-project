package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hpaes/go-api-final-project/src/core/application/usecase"
	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUserController_Execute(t *testing.T) {
	tests := []struct {
		name         string
		input        interface{}
		output       *usecase.RegisterUserOutput
		expectedCode int
		err          error
	}{
		{
			name: "success",
			input: usecase.RegisterUserInput{
				Id:    "1",
				Name:  "John Doe",
				Email: "johnDoe@email.com",
				Age:   25,
			},
			output: &usecase.RegisterUserOutput{
				Id:    1,
				Name:  "John Doe",
				Email: "johnDoe@email.com",
				Age:   25,
			},
			expectedCode: http.StatusCreated,
			err:          nil,
		},
		{
			name: "usecase error",
			input: usecase.RegisterUserInput{
				Id:    "1",
				Name:  "John Doe",
				Email: "johnDoe@email.com",
				Age:   25,
			},
			output:       nil,
			expectedCode: http.StatusInternalServerError,
			err:          assert.AnError,
		},
		{
			name:         "invalid request payload",
			input:        "invalid",
			output:       nil,
			expectedCode: http.StatusBadRequest,
			err:          custom_errors.NewInvalidRequestPayloadError("invalid request payload"),
		},
		{
			name: "user already exists",
			input: usecase.RegisterUserInput{
				Id:    "1",
				Name:  "John Doe",
				Email: "johnDoe@email.com",
				Age:   25,
			},
			output:       nil,
			expectedCode: http.StatusConflict,
			err:          custom_errors.NewUserAlreadyExistsError("johnDoe@email.com"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRU := &mockRegisterUser{}
			controller := NewRegisterUserController(mockRU)

			if tc.name != "invalid request payload" {
				mockRU.On("Execute", mock.Anything, mock.AnythingOfType("*usecase.RegisterUserInput")).Return(tc.output, tc.err)
			}

			w := httptest.NewRecorder()

			var body []byte
			if tc.name == "invalid request payload" {
				body = []byte(tc.input.(string))
			} else {
				body, _ = json.Marshal(tc.input)
			}
			req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))

			controller.Execute(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)
			mockRU.AssertExpectations(t)
		})
	}
}
