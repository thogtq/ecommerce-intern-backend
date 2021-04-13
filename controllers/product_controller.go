package controllers

import (
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/dao"
	"github.com/thogtq/ecommerce-server/errors"
	"github.com/thogtq/ecommerce-server/models"
	"github.com/thogtq/ecommerce-server/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var productDAO dao.ProductDAO

func CreateProduct(c *gin.Context) {
	productData := &models.Product{}
	c.BindJSON(productData)
	productData.ParentCategories = services.GetParentCategories(productData.Categories)
	productData.CreatedAt = time.Now()
	res, err := productDAO.InsertProduct(c.Request.Context(), productData)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, SuccessResponse(gin.H{"productID": res}))
	//Fix me
	//Handle empty images
}
func UploadProductImage(c *gin.Context) {
	file, err := c.FormFile("productImage")
	if err != nil {
		c.Error(errors.ErrNoFile)
		return
	}
	file.Filename = strings.ToLower(file.Filename)
	fileExt := path.Ext(file.Filename)
	acceptedExt := []string{".jpg", ".png", ".jpeg", ".gif"}
	if !contains(acceptedExt, fileExt) {
		c.Error(errors.ErrInvalidExtension)
		return
	}
	fileName, err := productDAO.UploadImage(c, file)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, SuccessResponse(gin.H{"fileName": fileName}))
	//Fix me
	//Response URL
}
func GetProducts(c *gin.Context) {
	filters := &models.ProductFilters{
		Sort:     c.Request.URL.Query().Get("sort"),
		Search:   c.Request.URL.Query().Get("search"),
		Category: c.Request.URL.Query().Get("category"),
	}
	products, err := productDAO.GetProducts(c, filters)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, SuccessResponse(gin.H{"products": products}))
}
func GetProduct(c *gin.Context) {

	productID := c.Request.URL.Query().Get("productID")
	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		//Fix me
		c.Error(nil)
		return
	}
	products, err := productDAO.GetProductByID(c, objectID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, SuccessResponse(gin.H{"product": products}))
}
