package errors

type ResourceNotFoundError struct {
	body string
}

func NewResourceNotFoundError(body string) error {
	return &ResourceNotFoundError{body: body}
}

func (e *ResourceNotFoundError) Error() string {
	return e.body
}
