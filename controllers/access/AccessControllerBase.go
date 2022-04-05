package access

import (
	"net/http"
	"trackingApp/models"
)

type AccessControllerBase interface {
	Post(w http.ResponseWriter, req *http.Request)
	Validate(req *http.Request) (models.Access,error)
}