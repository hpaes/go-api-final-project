package custom_errors

type UserIdIsRequiredError struct {
}

func NewUserIdIsRequiredError() *UserIdIsRequiredError {
	return &UserIdIsRequiredError{}
}

func (e *UserIdIsRequiredError) Error() string {
	return "User Id is required"
}
