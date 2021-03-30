package main

import (
	// "context"
	// "fmt"

	"context"
	"fmt"

	"github.com/subosito/gotenv"
	"github.com/thogtq/ecommerce-server/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbClient *mongo.Client

func init() {
	gotenv.Load()
}
func main() {
	dbClient = database.ConnectDatabase()
	defer database.DisconnectDatabase(dbClient)
	findFilter := bson.M{}
	res, err := dbClient.Database("ecommerce").Collection("test").Find(context.Background(), findFilter)
	_=err
	fmt.Printf("%v",res)
}
