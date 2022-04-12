package customError

import (
)

type ErrorSaving struct {
    ErrorMessage string
}

func (e *ErrorSaving) Error() string {
    return e.ErrorMessage
}