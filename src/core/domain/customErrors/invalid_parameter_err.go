package custom_errors

type InvalidParameterError struct {
	message   string
	parameter string
}

func NewInvalidParameterError(parameter string, message ...string) *InvalidParameterError {
	defaultMessage := "Invalid parameter"
	if len(message) > 0 {
		return &InvalidParameterError{parameter: parameter, message: message[0]}
	}
	return &InvalidParameterError{parameter: parameter, message: defaultMessage}
}

func (e *InvalidParameterError) Error() string {
	return e.message
}
