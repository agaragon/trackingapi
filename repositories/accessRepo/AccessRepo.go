package accessRepo

import (
	. "trackingApp/shared/dal"
	"os"
	. "trackingApp/models"
)

type AccessRepo struct {}

func (ar *AccessRepo) Save(access Access) bool {
	var db DbBase = &DbMongo{os.Getenv("DBURI")}
	// var db DbBase = &DbMongo{os.Getenv("mongodb://user:pass@mongodb:27017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false")}
	err := db.Save(access,"access")
	if err != nil{
		return false
	}
	return true
}