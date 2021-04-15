package dao

import (
	"context"
	"mime/multipart"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/constants"
	"github.com/thogtq/ecommerce-server/database"
	"github.com/thogtq/ecommerce-server/errors"
	"github.com/thogtq/ecommerce-server/helpers"
	"github.com/thogtq/ecommerce-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductDAO struct {
	productCollection   *mongo.Collection
	productImageTempDir string
	productImageDir     string
}

func (pd *ProductDAO) Init() {
	pd.productCollection = database.DBClient.Database("ecommerce").Collection("products")
}

func (pd *ProductDAO) GetProducts(c context.Context, filter *models.ProductFilters) (*[]models.Product, error) {
	pd.Init()
	productArray := []models.Product{}
	findFilter := bson.D{}
	findOptions := options.Find()
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
		findFilter = append(findFilter, bson.E{Key: "name", Value: primitive.Regex{Pattern: filter.Search, Options: ""}})
	}
	//Continue to implements color size brand price and available
	cur, err := pd.productCollection.Find(c, findFilter, findOptions)
	if err != nil {
		return nil, errors.ErrInternal(err.Error())
	}
	for cur.Next(c) {
		product := &models.Product{}
		err := cur.Decode(product)
		if err != nil {
			return nil, errors.ErrInternal(err.Error())
		}
		productArray = append(productArray, *product)
	}
	return &productArray, nil
}

func (pd *ProductDAO) GetProductByID(c context.Context, productID primitive.ObjectID) (*models.Product, error) {
	pd.Init()
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
	pd.Init()
	result, err := pd.productCollection.InsertOne(c, productObject)
	if err != nil {
		return "", errors.ErrInternal(err.Error())
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (pd *ProductDAO) UploadImage(c *gin.Context, file *multipart.FileHeader) (string, error) {
	pd.Init()
	file.Filename = helpers.GenerateUUID() + path.Ext(file.Filename)
	err := c.SaveUploadedFile(file, constants.PRODUCT_IMAGE_TEMP_DIR+file.Filename)
	if err != nil {
		return "", errors.ErrInternal(err.Error())
	}
	return file.Filename, nil
}
