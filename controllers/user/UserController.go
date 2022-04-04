package user

import (
	"net/http"
	"fmt"
)

type UserController struct {}

func (uc UserController) Post(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w,"Hello There 2")
}