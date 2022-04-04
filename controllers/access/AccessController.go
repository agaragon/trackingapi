package access

import (
	"fmt"
	"net/http"
)

type AccessController struct {}

func (ac AccessController) Post(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w,"Hello There 3")
}