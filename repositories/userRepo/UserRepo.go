package userRepo

import (
	. "trackingApp/shared/dal"
	. "trackingApp/models"
	. "trackingApp/shared/customError"
	. "trackingApp/shared/logger"
	"trackingApp/models"
	"encoding/json"
	"strings"
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


func (ur *UserRepo) FilterById(id []byte) models.User{
	var user models.User
	result,err := json.Marshal(ur.Db.FilterById(id,"user"))
	resultString := strings.Replace(string(result),"_id","clientId",-1)
	if err != nil{
		LogError(err)
	}
	err = json.Unmarshal([]byte(resultString),&user)
	if err != nil{
		LogError(err)
	}
	return user
}

