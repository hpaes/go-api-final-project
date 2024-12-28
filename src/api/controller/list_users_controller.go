package controller

import (
	"net/http"
	"strconv"

	"github.com/hpaes/go-api-final-project/src/api/handler"
	"github.com/hpaes/go-api-final-project/src/api/response"
	"github.com/hpaes/go-api-final-project/src/core/application/usecase"
	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
)

type ListUsersController struct {
	uc usecase.ListUsers
}

func NewListUsersController(uc usecase.ListUsers) *ListUsersController {
	return &ListUsersController{
		uc: uc,
	}
}

func (lu *ListUsersController) Execute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		handler.HandleError(w, custom_errors.NewInvalidRequestPayloadError(err.Error()))
		return
	}
	output, err := lu.uc.Execute(ctx, &usecase.ListUsersInput{Page: pageInt})
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	response.NewSuccessResponse(http.StatusOK, output).Send(w)
}
