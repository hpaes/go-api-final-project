package controller

import (
	"net/http"
	"strconv"

	"github.com/hpaes/go-api-final-project/src/api/handler"
	"github.com/hpaes/go-api-final-project/src/api/response"
	"github.com/hpaes/go-api-final-project/src/core/application/usecase"
	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
)

type GetUserDetailController struct {
	uc usecase.GetUserDetail
}

func NewGetUserDetailController(uc usecase.GetUserDetail) *GetUserDetailController {
	return &GetUserDetailController{
		uc: uc,
	}
}

func (gd *GetUserDetailController) Execute(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")

	if userID == "" {
		handler.HandleError(w, custom_errors.NewUserIdIsRequiredError())
		return
	}

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		handler.HandleError(w, custom_errors.NewInvalidRequestPayloadError(err.Error()))
		return
	}
	ctx := r.Context()
	output, err := gd.uc.Execute(ctx, &usecase.GetUserDetailInput{Id: userIDInt})
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	response.NewSuccessResponse(http.StatusOK, output).Send(w)
}
