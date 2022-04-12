package userRepo

import (
	"trackingApp/models"
)


type UserRepoBase interface {
	Save(user models.User) error
	// GetByClientId(clientId string) models.User
}