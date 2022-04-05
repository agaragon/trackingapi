package userRepo

import (
	"trackingApp/models"
	. "trackingApp/shared/dal"
)


type UserRepoBase interface {
	Db DbBase
	Save(user models.User) bool
	GetByClientId(clientId string) models.User
}