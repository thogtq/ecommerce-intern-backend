package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/thogtq/ecommerce-server/database"
	"go.mongodb.org/mongo-driver/bson"
)

func Test() string{
	r := ""
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := database.DBClient.Database("tests").Collection("user").Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		r = fmt.Sprint(result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return r
}
