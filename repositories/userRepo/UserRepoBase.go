package userRepo

import (
	"trackingApp/models"
)


type UserRepoBase interface {
	Save(user models.User) error
	FilterById(id []byte) models.User
}