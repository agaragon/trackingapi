package customError

type ToManyRowsFound struct {
    ErrorMessage string
}

func (e *ToManyRowsFound) Error() string {
    return e.ErrorMessage
}