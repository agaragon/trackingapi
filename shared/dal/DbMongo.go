package dal
import (
	. "trackingApp/shared/customError"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"os"
	. "trackingApp/shared/logger"
	"go.mongodb.org/mongo-driver/bson"
	"encoding/json"
)


type DbMongo struct {
	Uri string 
}

func (db *DbMongo) Save(object interface{},tableName string) error{
	clientOptions := options.Client().ApplyURI(db.Uri)
	dbConn, err := mongo.Connect(context.TODO(), clientOptions)
	if err !=  nil {
		LogError(err)
		return &DbConnectionError{"Failed to connect to database"}
	}
	collection := dbConn.Database(os.Getenv("DATABASE")).Collection(tableName)
	_,err = collection.InsertOne(context.TODO(), object)
	return err
}

func (db *DbMongo) FilterById(id []byte, tableName string) interface{}{
	clientOptions := options.Client().ApplyURI(db.Uri)
	dbConn, err := mongo.Connect(context.TODO(), clientOptions)
	if err !=  nil {
		LogError(err)
		return &DbConnectionError{"Failed to connect to database"}
	}
	collection := dbConn.Database(os.Getenv("DATABASE")).Collection(tableName)
	var filter interface{}
	err = json.Unmarshal(id,&filter)
	if err != nil{
		LogError(err)
	}
	var result bson.M
	err = collection.FindOne(context.TODO(),bson.M{"_id":&filter}).Decode(&result)
	if err != nil{
		LogError(err)
	}
	return result
}