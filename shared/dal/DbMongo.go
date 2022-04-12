package dal
import (
	. "trackingApp/shared/customError"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"os"
)


type DbMongo struct {
	Uri string 
}

func (db *DbMongo) Save(object interface{},tableName string) error{
	clientOptions := options.Client().ApplyURI(db.Uri)
	dbConn, err := mongo.Connect(context.TODO(), clientOptions)
	if err !=  nil {
		return &DbConnectionError{"Failed to connect to database"}
	}
	collection := dbConn.Database(os.Getenv("DATABASE")).Collection(tableName)
	_,err = collection.InsertOne(context.TODO(), object)
	return err
}