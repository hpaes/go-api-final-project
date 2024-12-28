package custom_errors

import "fmt"

type InvalidRequestPayloadError struct {
	message string
}

func NewInvalidRequestPayloadError(message string) *InvalidRequestPayloadError {
	return &InvalidRequestPayloadError{message: message}
}

func (e *InvalidRequestPayloadError) Error() string {
	return fmt.Sprintf("Invalid request payload: %s", e.message)

}
