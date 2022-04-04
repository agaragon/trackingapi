package access

import (
	"net/http"
)

type AccessControllerBase interface {
	Post(w http.ResponseWriter, req *http.Request)
}