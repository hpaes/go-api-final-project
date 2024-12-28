package handler

import (
	"net/http"

	"github.com/hpaes/go-api-final-project/src/api/response"
	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
)

func HandleError(w http.ResponseWriter, err error) {
	var status int

	switch err.(type) {
	case *custom_errors.ParameterRequiredError,
		*custom_errors.InvalidParameterError,
		*custom_errors.InvalidRequestPayloadError,
		*custom_errors.UserIdIsRequiredError:
		status = http.StatusBadRequest
	case *custom_errors.UserNotFoundError:
		status = http.StatusNotFound
	case *custom_errors.UserAlreadyExistsError:
		status = http.StatusConflict
	default:
		status = http.StatusInternalServerError
	}

	response.NewErrorResponse(status, err).Send(w)
}
