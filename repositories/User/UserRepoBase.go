package repositories

import (
	"trackingApp/models"
)


type UserRepoBase interface {
	Save() bool
	GetByClientId(clientId string) models.User
}