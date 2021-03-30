package database

import (
	"context"
	"log"
	// "os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDatabase() *mongo.Client {
	// var (
	// 	HOST = os.Getenv("MONGODB_HOST")
	// 	PORT = os.Getenv("MONGODB_PORT")
	// )
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://172.17.0.2:27017"))
	if err != nil {
		log.Panicf("Can not connect to database : %v", err)
	}
	return client
}
func DisconnectDatabase(dbClient *mongo.Client) {
	if err := dbClient.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
