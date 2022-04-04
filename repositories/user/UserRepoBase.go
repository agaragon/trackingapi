package user

import (
	"trackingApp/models"
)


type UserRepoBase interface {
	Save() bool
	GetByClientId(clientId string) models.User
}