package userRepo

import (
	"trackingApp/models"
)


type UserRepoBase interface {
	Save(user models.User) error
	Filter(id []byte, tableName string) interface{}
	// GetByClientId(clientId string) models.User
}