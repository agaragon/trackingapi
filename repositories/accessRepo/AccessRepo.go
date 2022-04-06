package accessRepo

import (
	. "trackingApp/shared/dal"
	. "trackingApp/models"
	"fmt"
	"os"
)

type AccessRepo struct {
	Db DbBase
}

func (ar *AccessRepo) Save(access Access) bool {
	
	err := ar.Db.Save(access,"access")
	if err != nil{
		fmt.Println(os.Getenv("DATABASE_URL"))
		fmt.Println(err)
		return false
	}
	return true
}