package models

import (
	"context"
	"fmt"
	"time"

	"github.com/thogtq/ecommerce-server/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection

type Product struct {
	ProductID      primitive.ObjectID `bson:"_id,omitempty" json:"productID"`
	Name           string             `bson:"name" json:"name"`
	AddDate        time.Time          `bson:"addDate" json:"addDate"`
	Categories     []string           `bson:"categories" json:"categories"`
	ParentCategory string             `bson:"parentCategory" json:"parentCategory"`
	Images         []string           `bson:"images" json:"images"`
	Brand          string             `bson:"brand" json:"brand"`
	Price          float32            `bson:"price" json:"price"`
	Sizes          []string           `bson:"size" json:"size"`
	Colors         []string           `bson:"color" json:"color"`
	Quantity       int                `bson:"quantity" json:"quantity"`
	Sold           int                `bson:"sold" json:"sold"`
	Description    string             `bson:"description" json:"description"`
}

func (Product) GetProductsByCategory(categoryName string) (*[]Product, error) {
	productArray := []Product{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	productCollection = database.DBClient.Database("ecommerce").Collection("products")
	_ = ctx
	findFilter := bson.M{"$or": bson.M{"categories": categoryName, "parentCategory": categoryName}}
	cur, err := productCollection.Find(ctx, findFilter)
	if err != nil {
		return nil, fmt.Errorf("can not get product:%v", err)
	}
	for cur.Next(ctx) {
		product := &Product{}
		err := cur.Decode(product)
		if err != nil {
			return nil, fmt.Errorf("can not decode data:%v", err)
		}
		productArray = append(productArray, *product)
	}
	return &productArray, nil
}
func (Product) GetProductByID(productID string) (*Product, error) {
	return nil, nil
}
func (Product) InsertProduct(productObject *Product) (string, error) {
	productObject = &Product{
		Name:           "Test product",
		AddDate:        time.Now(),
		Categories:     []string{"cate1", "cate2"},
		ParentCategory: "men",
		Images:         []string{"adf.jpg", "sdf.jpg", "sdf.jpg"},
		Brand:          "Zara",
		Price:          99,
		Sizes:          []string{"X", "XL"},
		Colors:         []string{"Red", "Green"},
		Quantity:       9,
		Sold:           2,
		Description:    "Desc",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	productCollection = database.DBClient.Database("ecommerce").Collection("products")
	result, err := productCollection.InsertOne(ctx, productObject)
	if err != nil {
		return "", fmt.Errorf("product not inserted:%v", err)
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
