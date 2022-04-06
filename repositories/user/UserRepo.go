package userRepo

import (
	. "trackingApp/shared/dal"
	"trackingApp/models"
	"fmt"
)

type UserRepo struct {
	Db DbBase 
}

func (ur *UserRepo) Save(user models.User) bool {
	err := ar.Db.Save(user,"user")
	if err != nil{
		fmt.Println(os.Getenv("DATABASE_URL"))
		fmt.Println(err)
		return false
	}
}