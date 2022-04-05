package dal
import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
)


type DbMongo struct {
	Uri string 
}

func (db *DbMongo) StartConnection()(*mongo.Client,error){
	clientOptions :=options.Client().ApplyURI(db.Uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	return client,err
}