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

type RegisterUserController struct {
	ru usecase.RegisterUser
}

func NewRegisterUserController(uc usecase.RegisterUser) *RegisterUserController {
	return &RegisterUserController{
		ru: uc,
	}
}

func (rc *RegisterUserController) Execute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonBody, err := io.ReadAll(r.Body)
	if err != nil {
		handler.HandleError(w, custom_errors.NewInvalidRequestPayloadError(err.Error()))
		return
	}
	var input usecase.RegisterUserInput
	if err := json.Unmarshal(jsonBody, &input); err != nil {
		handler.HandleError(w, custom_errors.NewInvalidRequestPayloadError(err.Error()))
		return
	}
	output, err := rc.ru.Execute(ctx, &input)
	if err != nil {
		handler.HandleError(w, err)
		return
	}
	response.NewSuccessResponse(http.StatusCreated, output).Send(w)
}
