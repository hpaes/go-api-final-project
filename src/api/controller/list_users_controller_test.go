package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hpaes/go-api-final-project/src/core/application/usecase"
	"github.com/hpaes/go-api-final-project/src/core/domain"
	"github.com/hpaes/go-api-final-project/src/core/domain/vo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListUsersController_Execute(t *testing.T) {
	email, _ := vo.NewEmail("johnDoe@email.com")
	output := []domain.User{
		{
			Id:    1,
			Name:  "John Doe",
			Email: *email,
			Age:   25,
		},
	}

	tests := []struct {
		name         string
		page         string
		output       []domain.User
		expectedCode int
		err          error
	}{
		{
			name:         "success",
			page:         "1",
			output:       output,
			expectedCode: http.StatusOK,
			err:          nil,
		},
		{
			name:         "usecase error",
			page:         "1",
			output:       nil,
			expectedCode: http.StatusInternalServerError,
			err:          assert.AnError,
		},
		{
			name:         "invalid page",
			page:         "invalid",
			output:       nil,
			expectedCode: http.StatusBadRequest,
			err:          assert.AnError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockUC := &mockListUsers{}
			controller := NewListUsersController(mockUC)

			if tc.name == "success" || tc.name == "usecase error" {
				mockUC.On("Execute", mock.Anything, &usecase.ListUsersInput{Page: 1}).Return(tc.output, tc.err)
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/users?page="+tc.page, nil)

			controller.Execute(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)
			mockUC.AssertExpectations(t)
		})
	}
}
