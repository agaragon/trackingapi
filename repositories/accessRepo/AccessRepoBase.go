package accessRepo

import (
	"trackingApp/models"
)


type AccessRepoBase interface {
	Save(access models.Access) error
}