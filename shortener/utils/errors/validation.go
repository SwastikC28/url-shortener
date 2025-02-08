package errors

type ValidationError struct {
	body string
}

func NewValidationError(body string) error {
	return &ValidationError{
		body: body,
	}
}

func (e *ValidationError) Error() string {
	return e.body
}
