package controller

import (
	"net/http"
	"strconv"

	"github.com/hpaes/go-api-final-project/src/api/handler"
	"github.com/hpaes/go-api-final-project/src/api/response"
	"github.com/hpaes/go-api-final-project/src/core/application/usecase"
	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
)

type RemoveUserController struct {
	uc usecase.RemoveUser
}

func NewRemoveUserController(uc usecase.RemoveUser) *RemoveUserController {
	return &RemoveUserController{
		uc: uc,
	}
}

func (ru *RemoveUserController) Execute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
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
	err = ru.uc.Execute(ctx, &usecase.RemoveUserInput{Id: userIDInt})
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	response.NewSuccessResponse(http.StatusOK, map[string]string{"message": "User removed successfully"}).Send(w)
}
