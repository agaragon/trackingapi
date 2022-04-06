package accessRepo

import (
	. "trackingApp/shared/dal"
	. "trackingApp/models"
)

type AccessRepo struct {
	Db DbBase
}

func (ar *AccessRepo) Save(access Access) bool {
	
	err := ar.Db.Save(access,"access")
	if err != nil{
		return false
	}
	return true
}