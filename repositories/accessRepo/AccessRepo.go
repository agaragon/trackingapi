package accessRepo

import (
	. "trackingApp/shared/dal"
	"log"
	"context"
	. "trackingApp/models"
)

type AccessRepo struct {}

func (ar *AccessRepo) Save(access Access) bool {
	db := DbManager{"mongodb://user:pass@mongodb:27017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false"}
	dbConn,err := db.StartConnection()
	if err !=  nil {
		log.Fatal(err)
	}
	collection := dbConn.Database("trackingapp").Collection("access")
	_, err = collection.InsertOne(context.TODO(), access)
	if err != nil{
		return false
	}
	return true
}