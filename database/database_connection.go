package database

import (
	"context"
	"log"

	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBClient *mongo.Client

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var (
		DEFAULT_HOST = os.Getenv("MONGODB_HOST")
		DEFAULT_PORT = os.Getenv("MONGODB_PORT")
	)
	if DEFAULT_HOST == "" || DEFAULT_PORT == "" {
		log.Panicf("unable to load DB CONFIG variables")
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DEFAULT_HOST+`:`+DEFAULT_PORT))
	if err != nil {
		log.Panicf("can not connect to database : %v", err)
	}
	DBClient = client
}
func Disconnect() {
	if err := DBClient.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
