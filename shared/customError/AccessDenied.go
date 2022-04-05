package customError

import (
)

type AccessDenied struct {
    ErrorMessage string
}

func (e *AccessDenied) Error() string {
    return e.ErrorMessage
}