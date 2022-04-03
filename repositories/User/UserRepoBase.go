package repositories

import (
	"trackingApp/models"
)


type UserRepoBase interface {
	Save() bool
	Get() models.User
}