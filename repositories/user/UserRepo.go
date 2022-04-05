package userRepo

import (
	. "trackingApp/shared/dal"
	"trackingApp/models"
	
)

type UserRepo struct {
	Db DbBase 
}

func (ur *UserRepo) Save() bool {
	dbConn,err := db.StartConnection()
	if err !=  nil {
		log.Fatal(err)
	}
	collection := dbConn.Database("trackingapp").Collection("access")
	insertResult, err := collection.InsertOne(context.TODO(), access)
	return insertResult,err
}