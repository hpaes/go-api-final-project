package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/hpaes/go-api-final-project/src/api/handler"
	"github.com/hpaes/go-api-final-project/src/api/response"
	"github.com/hpaes/go-api-final-project/src/core/application/usecase"
	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
)

type UpdateUserController struct {
	uc usecase.UpdateUser
}

func NewUpdateUserController(uc usecase.UpdateUser) *UpdateUserController {
	return &UpdateUserController{uc: uc}
}

func (uu *UpdateUserController) Execute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonBody, err := io.ReadAll(r.Body)
	if err != nil {
		handler.HandleError(w, custom_errors.NewInvalidRequestPayloadError(err.Error()))
		return
	}
	var input usecase.UpdateUserInput
	if err := json.Unmarshal(jsonBody, &input); err != nil {
		handler.HandleError(w, custom_errors.NewInvalidRequestPayloadError(err.Error()))
		return
	}
	output, err := uu.uc.Execute(ctx, &input)
	if err != nil {
		handler.HandleError(w, err)
		return
	}
	response.NewSuccessResponse(http.StatusOK, output).Send(w)
}
