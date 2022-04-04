package user

import (
	"net/http"
)

type UserControllerBase interface{
	Post(w http.ResponseWriter, req *http.Request)
}