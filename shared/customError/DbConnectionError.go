package customError

import (
)

type DbConnectionError struct {
    ErrorMessage string
}

func (e *DbConnectionError) Error() string {
    return e.ErrorMessage
}