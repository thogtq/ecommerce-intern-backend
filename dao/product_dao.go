package dao

import (
	"context"
	"encoding/json"
	"math"
	"mime/multipart"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/constants"
	"github.com/thogtq/ecommerce-server/database"
	"github.com/thogtq/ecommerce-server/errors"
	"github.com/thogtq/ecommerce-server/models"
	"github.com/thogtq/ecommerce-server/pkg/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductDAO struct {
	productCollection *mongo.Collection
}

func (pd *ProductDAO) New() *ProductDAO {
	pd.productCollection = database.DBClient.Database("ecommerce").Collection("products")
	return pd
}

func (pd *ProductDAO) GetProducts(c context.Context, filter *models.ProductFilters) (*[]models.Product, int64, error) {
	productArray := []models.Product{}
	counts := int64(0)
	rowSkip := (filter.Page - 1) * filter.Limit
	findFilter := bson.D{}
	findOptions := options.Find()
	findOptions.SetLimit(int64(filter.Limit))
	findOptions.SetSkip(int64(rowSkip))
	//Sort options
	if filter.SortBy != "" {
		findOptions.SetSort(bson.D{{Key: filter.SortBy, Value: filter.SortOrder}})
	}
	//Filter attribute
	if filter.Category != "" {
		key := "categories"
		if !strings.Contains(filter.Category, "/") {
			key = "parentCategories"
		}
		findFilter = append(findFilter, bson.E{Key: key, Value: filter.Category})
	}
	//Search by name
	if filter.Search != "" {
		findFilter = append(findFilter, bson.E{Key: "name", Value: primitive.Regex{Pattern: filter.Search, Options: "i"}})
	}
	//Filter by attributes
	if filter.Color != "" {
		findFilter = append(findFilter, bson.E{Key: "colors", Value: filter.Color})
	}
	if filter.Size != "" {
		findFilter = append(findFilter, bson.E{Key: "sizes", Value: filter.Size})
	}
	if filter.Brand != "" {
		findFilter = append(findFilter, bson.E{Key: "brand", Value: filter.Brand})
	}
	if filter.MaxPrice != 0 {
		findFilter = append(findFilter, bson.E{Key: "price", Value: bson.D{
			{Key: "$gte", Value: filter.MinPrice},
			{Key: "$lte", Value: filter.MaxPrice},
		}})
	}
	// if filter.Available != "" {
	// 	//Fix me
	// 	//Not working
	// 	operator := "$ne"
	// 	if filter.Available == "out" {
	// 		operator = "$eq"
	// 	}
	// 	findFilter = append(findFilter, bson.E{Key: "$expr", Value: bson.E{
	// 		Key: operator, Value: []string{"$sold", "$quantity"},
	// 	}})
	// }
	cur, err := pd.productCollection.Find(c, findFilter, findOptions)
	if err != nil {
		return nil, 0, errors.ErrInternal(err.Error())
	}
	counts, _ = pd.productCollection.CountDocuments(c, findFilter)
	for cur.Next(c) {
		product := &models.Product{}
		err := cur.Decode(product)
		if err != nil {
			return nil, 0, errors.ErrInternal(err.Error())
		}
		productArray = append(productArray, *product)
	}
	return &productArray, counts, nil
}

func (pd *ProductDAO) GetProductByID(c context.Context, productID primitive.ObjectID) (*models.Product, error) {
	product := &models.Product{}
	filter := bson.M{"_id": productID}
	result := pd.productCollection.FindOne(c, filter)
	if result.Err() == mongo.ErrNoDocuments {
		return nil, errors.ErrProductNotFound
	}
	if result.Err() != nil {
		return nil, errors.ErrInternal(result.Err().Error())
	}
	if err := result.Decode(product); err != nil {
		return nil, errors.ErrInternal(err.Error())
	}
	return product, nil
}

func (pd *ProductDAO) InsertProduct(c context.Context, productObject *models.Product) (string, error) {
	result, err := pd.productCollection.InsertOne(c, productObject)
	if err != nil {
		return "", errors.ErrInternal(err.Error())
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (pd *ProductDAO) UploadImage(c *gin.Context, file *multipart.FileHeader) (string, error) {
	file.Filename = uuid.NewShortUUID() + path.Ext(file.Filename)
	err := c.SaveUploadedFile(file, constants.PRODUCT_IMAGE_TEMP_DIR+file.Filename)
	if err != nil {
		return "", errors.ErrInternal(err.Error())
	}
	return file.Filename, nil
}
func (pd *ProductDAO) DeleteProduct(c context.Context, productID primitive.ObjectID) error {
	_, err := pd.productCollection.DeleteOne(c, bson.D{{Key: "_id", Value: productID}})
	if err != nil {
		return errors.ErrProductNotFound
	}
	return nil
}
func (pd *ProductDAO) UpdateProduct(c context.Context, product *models.Product) error {
	objectID, err := primitive.ObjectIDFromHex(product.ProductID.Hex())
	if err != nil {
		return errors.ErrProductNotFound
	}
	filter := bson.D{{Key: "_id", Value: objectID}}
	update := bson.M{}

	pBytes, _ := json.Marshal(product)
	json.Unmarshal(pBytes, &update)
	delete(update, "sold")
	delete(update, "profit")
	delete(update, "reviews")
	delete(update, "star")
	result, err := pd.productCollection.UpdateOne(c, filter, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return errors.ErrInternal(err.Error())
	}
	if result.MatchedCount == 0 {
		return errors.ErrProductNotFound
	}
	return nil
}
func (pd *ProductDAO) UpdateProductSale(c context.Context, productID string, sold int, profit int) error {
	ObjectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return errors.ErrProductNotFound
	}
	update := bson.D{{"$inc", bson.D{
		{"profit", profit},
		{"sold", sold},
	}}}
	result, err := pd.productCollection.UpdateOne(c, bson.M{"_id": ObjectID}, update)
	if err != nil {
		return errors.ErrInternal(err.Error())
	}
	if result.MatchedCount == 0 {
		return errors.ErrProductNotFound
	}
	return nil
}
func (pd *ProductDAO) UpdateProductReview(c context.Context, productID string, star int, review int) error {
	var (
		productReviews int
		productStar    int
		update         bson.D
		newStar        int
		newReviews     int
	)
	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return errors.ErrProductNotFound
	}
	product, err := pd.GetProductByID(c, objectID)
	if err != nil {
		return err
	}
	productReviews = product.Reviews
	productStar = product.Star

	floatNewStar := float64((productStar*productReviews + star)) / float64((productReviews + review))
	newStar = int(math.Round(floatNewStar))
	newReviews = productReviews + review
	update = bson.D{{"$set", bson.D{
		{"star", newStar},
		{"reviews", newReviews},
	}}}
	result, err := pd.productCollection.UpdateOne(c, bson.M{"_id": objectID}, update)
	if err != nil {
		return errors.ErrInternal(err.Error())
	}
	if result.MatchedCount == 0 {
		return errors.ErrProductNotFound
	}
	return nil
}
