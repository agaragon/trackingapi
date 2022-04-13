package userRepo

import (
	. "trackingApp/shared/dal"
	. "trackingApp/models"
	. "trackingApp/shared/customError"
	. "trackingApp/shared/logger"
)

type UserRepo struct {
	Db DbBase 
}


func (ur *UserRepo) Save(user User) error {
	err := ur.Db.Save(user,"user")
	if err != nil{
		LogError(err)
		return &ErrorSaving{"Unable to save to database"}
	}
	return nil
}


func (ur *UserRepo) Filter(id []byte, tableName string) interface{}{
	return ur.Db.Filter(id,tableName)
}

// func (ur *UserRepo)GetByClientId(clientId string) User{

// }
