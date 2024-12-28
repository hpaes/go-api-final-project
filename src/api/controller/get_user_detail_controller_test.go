package controller

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/hpaes/go-api-final-project/src/core/application/usecase"
	"github.com/hpaes/go-api-final-project/src/core/domain"
	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
	"github.com/hpaes/go-api-final-project/src/core/domain/vo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserDetailController_Execute(t *testing.T) {
	email, _ := vo.NewEmail("johnDoe@email.com")
	output := domain.User{
		Id:    1,
		Name:  "John Doe",
		Email: *email,
		Age:   25,
	}

	tests := []struct {
		name         string
		userID       string
		output       domain.User
		expectedCode int
		err          error
	}{
		{
			name:         "success",
			userID:       "1",
			output:       output,
			expectedCode: http.StatusOK,
			err:          nil,
		},
		{
			name:         "usecase error",
			userID:       "1",
			output:       domain.User{},
			expectedCode: http.StatusInternalServerError,
			err:          assert.AnError,
		},
		{
			name:         "invalid user id",
			userID:       "invalid",
			output:       domain.User{},
			expectedCode: http.StatusBadRequest,
			err:          custom_errors.NewInvalidRequestPayloadError("invalid user id"),
		},
		{
			name:         "user id is required",
			userID:       "",
			output:       domain.User{},
			expectedCode: http.StatusBadRequest,
			err:          custom_errors.NewUserIdIsRequiredError(),
		},
		{
			name:         "user not found",
			userID:       "1",
			output:       domain.User{},
			expectedCode: http.StatusNotFound,
			err:          custom_errors.NewUserNotFoundError("User", "1"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockUC := &mockGetUserDetail{}
			controller := NewGetUserDetailController(mockUC)

			if tc.name == "success" || tc.name == "usecase error" || tc.name == "user not found" {
				userIDInt, _ := strconv.Atoi(tc.userID)
				mockUC.On("Execute", mock.Anything, &usecase.GetUserDetailInput{Id: userIDInt}).Return(tc.output, tc.err)
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/user?userId="+tc.userID, nil)

			controller.Execute(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)
			mockUC.AssertExpectations(t)
		})
	}
}
