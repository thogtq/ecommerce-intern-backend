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

func (pd *OrderDAO) New() *OrderDAO {
	pd.orderCollection = database.DBClient.Database("ecommerce").Collection("orders")
	return pd
}

//Fix me
func (pd *OrderDAO) Init() {
	pd.orderCollection = database.DBClient.Database("ecommerce").Collection("orders")
}
func (pd *OrderDAO) InsertOrder(c context.Context, order *models.Order) (string, error) {
	pd.Init()
	result, err := pd.orderCollection.InsertOne(c, order)
	if err != nil {
		return "", nil
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
func (pd *OrderDAO) GetOrders(c context.Context, filter *models.OrderFilter) (*[]models.Order, int64, error) {
	pd.Init()
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
	cur, err := pd.orderCollection.Find(c, findFilter, findOptions)
	if err != nil {
		return nil, 0, errors.ErrInternal(err.Error())
	}
	counts, _ = pd.orderCollection.CountDocuments(c, findFilter)
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
