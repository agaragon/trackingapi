package dal
import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"os"
	"log"
)


type DbMongo struct {
	Uri string 
}

func (db *DbMongo) Save(object interface{},tableName string) error{
	clientOptions := options.Client().ApplyURI(db.Uri)
	dbConn, err := mongo.Connect(context.TODO(), clientOptions)
	if err !=  nil {
		log.Fatal(err)
	}
	collection := dbConn.Database(os.Getenv("DATABASE")).Collection(tableName)
	_,err = collection.InsertOne(context.TODO(), object)
	return err
}