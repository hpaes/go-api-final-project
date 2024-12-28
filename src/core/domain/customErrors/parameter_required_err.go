package custom_errors

import "fmt"

type ParameterRequiredError struct {
	message   string
	parameter string
}

func NewParameterRequiredError(parameter string) *ParameterRequiredError {
	return &ParameterRequiredError{parameter: parameter}
}

func (e *ParameterRequiredError) Error() string {
	return fmt.Sprintf("Parameter %s is required", e.parameter)
}
