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

func TestUpdateUserController_Execute(t *testing.T) {
	tests := []struct {
		name         string
		input        interface{}
		output       *usecase.UpdateUserOutput
		expectedCode int
		err          error
	}{
		{
			name: "success",
			input: usecase.UpdateUserInput{
				Id:    1,
				Name:  "John Doe",
				Email: "johnDoe@email.com",
				Age:   25,
			},
			output: &usecase.UpdateUserOutput{
				Id:    1,
				Name:  "John Doe",
				Email: "johnDoe@email.com",
				Age:   25,
			},
			expectedCode: http.StatusOK,
			err:          nil,
		},
		{
			name: "usecase error",
			input: usecase.UpdateUserInput{
				Id:    1,
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockUC := &mockUpdateUser{}
			controller := NewUpdateUserController(mockUC)

			if tc.name != "invalid request payload" {
				mockUC.On("Execute", mock.Anything, mock.AnythingOfType("*usecase.UpdateUserInput")).Return(tc.output, tc.err)
			}

			w := httptest.NewRecorder()

			var body []byte
			if tc.name == "invalid request payload" {
				body = []byte(tc.input.(string))
			} else {
				body, _ = json.Marshal(tc.input)
			}
			req := httptest.NewRequest(http.MethodPut, "/user", bytes.NewBuffer(body))

			controller.Execute(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)
			mockUC.AssertExpectations(t)
		})
	}
}
