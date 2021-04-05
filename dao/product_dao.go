package dao

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/database"
	"github.com/thogtq/ecommerce-server/helpers"
	"github.com/thogtq/ecommerce-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductDAO struct {
	productCollection   *mongo.Collection
	productImageTempDir string
	productImageDir     string
}

func (pd *ProductDAO) Init() {
	pd.productCollection = database.DBClient.Database("ecommerce").Collection("products")
	pd.productImageTempDir = "./files/images/temp/"
	pd.productImageDir = "./files/images/products/"
}

func (pd *ProductDAO) GetProductsByCategory(c context.Context, categoryName string) (*[]models.Product, error) {
	pd.Init()
	productArray := []models.Product{}
	findFilter := bson.M{"$or": bson.M{"categories": categoryName, "parentCategory": categoryName}}
	cur, err := pd.productCollection.Find(c, findFilter)
	if err != nil {
		//error
		return nil, fmt.Errorf("can not get product:%v", err)
	}
	for cur.Next(c) {
		product := &models.Product{}
		err := cur.Decode(product)
		if err != nil {
			//error
			return nil, fmt.Errorf("can not decode data:%v", err)
		}
		productArray = append(productArray, *product)
	}
	return &productArray, nil
}

func (pd *ProductDAO) GetProductByID(c context.Context, productID string) (*models.Product, error) {
	return nil, nil
}

func (pd *ProductDAO) InsertProduct(c context.Context, productObject *models.Product) (string, error) {
	pd.Init()
	result, err := pd.productCollection.InsertOne(c, productObject)
	if err != nil {
		//error
		return "", fmt.Errorf("product not inserted:%v", err)
	}
	for _, image := range productObject.Images {
		err := os.Rename(pd.productImageTempDir+image, pd.productImageDir+image)
		if err != nil {
			//error
			return "", fmt.Errorf("error when move temp image:%v", err)
		}
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (pd *ProductDAO) UploadImage(c *gin.Context, file *multipart.FileHeader) (string, error) {
	pd.Init()
	file.Filename = helpers.GenerateUUID() + path.Ext(file.Filename)
	err := c.SaveUploadedFile(file, pd.productImageTempDir+file.Filename)
	if err != nil {
		//error
		return "", err
	}
	return file.Filename, nil
}
