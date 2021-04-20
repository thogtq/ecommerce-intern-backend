package dao

import (
	"context"

	"github.com/thogtq/ecommerce-server/database"
	"github.com/thogtq/ecommerce-server/errors"
	"github.com/thogtq/ecommerce-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderDAO struct {
	orderCollection *mongo.Collection
}

func (od *OrderDAO) New() *OrderDAO {
	od.orderCollection = database.DBClient.Database("ecommerce").Collection("orders")
	return od
}

//Fix me
func (od *OrderDAO) Init() {
	od.orderCollection = database.DBClient.Database("ecommerce").Collection("orders")
}
func (od *OrderDAO) InsertOrder(c context.Context, order *models.Order) (string, error) {
	od.Init()
	result, err := od.orderCollection.InsertOne(c, order)
	if err != nil {
		return "", nil
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
func (od *OrderDAO) GetOrders(c context.Context, filter *models.OrderFilter) (*[]models.Order, int64, error) {
	od.Init()
	var (
		skipRows    = (filter.Page - 1) * filter.Limit
		findFilter  = bson.D{}
		findOptions = options.Find()
		orders      = []models.Order{}
		counts      int64
	)
	findOptions.SetSort(bson.D{{Key: "orderDate", Value: -1}})
	findOptions.SetSkip(int64(skipRows))
	findOptions.SetLimit(int64(filter.Limit))
	if filter.Search != "" {
		findFilter = append(findFilter, bson.E{
			Key:   "orderID",
			Value: primitive.Regex{Pattern: filter.Search, Options: ""},
		})
	}
	cur, err := od.orderCollection.Find(c, findFilter, findOptions)
	if err != nil {
		return nil, 0, errors.ErrInternal(err.Error())
	}
	counts, _ = od.orderCollection.CountDocuments(c, findFilter)
	for cur.Next(c) {
		order := &models.Order{}
		err := cur.Decode(order)
		if err != nil {
			return nil, 0, errors.ErrInternal(err.Error())
		}
		orders = append(orders, *order)
	}
	return &orders, counts, nil
}
func (od *OrderDAO) UpdateStatus(c context.Context, orderID, status string) error {
	update := bson.D{{"$set", bson.D{{"status", status}}}}
	result, err := od.orderCollection.UpdateOne(c, bson.M{"orderID": orderID}, update)
	if err != nil {
		return errors.ErrInternal(err.Error())
	}
	if result.MatchedCount == 0 {
		return errors.ErrOrderIDNotFound
	}
	return nil
}
