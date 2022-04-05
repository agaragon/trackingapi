package dal

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DbBase interface {
	StartConnection() (*mongo.Client,error)
}