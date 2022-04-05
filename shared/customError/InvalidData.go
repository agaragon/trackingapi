package customError

type InvalidData struct {
    ErrorMessage string
}

func (e *InvalidData) Error() string {
    return e.ErrorMessage
}