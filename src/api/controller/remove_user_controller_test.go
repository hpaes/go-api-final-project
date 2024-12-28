package controller

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/hpaes/go-api-final-project/src/core/application/usecase"
	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRemoveUserController_Execute(t *testing.T) {
	tests := []struct {
		name         string
		userID       string
		expectedCode int
		err          error
	}{
		{
			name:         "success",
			userID:       "1",
			expectedCode: http.StatusOK,
			err:          nil,
		},
		{
			name:         "usecase error",
			userID:       "1",
			expectedCode: http.StatusInternalServerError,
			err:          assert.AnError,
		},
		{
			name:         "invalid user id",
			userID:       "invalid",
			expectedCode: http.StatusBadRequest,
			err:          custom_errors.NewInvalidRequestPayloadError("invalid user id"),
		},
		{
			name:         "user id is required",
			userID:       "",
			expectedCode: http.StatusBadRequest,
			err:          custom_errors.NewUserIdIsRequiredError(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockUC := &mockRemoveUser{}
			controller := NewRemoveUserController(mockUC)

			if tc.name == "success" || tc.name == "usecase error" {
				userIDInt, _ := strconv.Atoi(tc.userID)
				mockUC.On("Execute", mock.Anything, &usecase.RemoveUserInput{Id: userIDInt}).Return(tc.err)
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, "/user?userId="+tc.userID, nil)

			controller.Execute(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)
			mockUC.AssertExpectations(t)
		})
	}
}
