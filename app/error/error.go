package applicationerror

import "fmt"

type ApplicationError struct {
	msg string
	err error
}

func NewApplicationError(msg string, err error) *ApplicationError {
	return &ApplicationError{msg, err}
}

func (e *ApplicationError) Error() string {
	return fmt.Sprintf("msg: %s err: %s\n", e.msg, e.err.Error())
}
