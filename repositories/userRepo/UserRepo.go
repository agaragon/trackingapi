package userRepo

import (
	. "trackingApp/shared/dal"
	. "trackingApp/models"
)

type UserRepo struct {
	Db DbBase 
}


func (ur *UserRepo) Save(user User) error {
	err := ur.Db.Save(user,"user")
	if err != nil{
		return err
	}
	return nil
}