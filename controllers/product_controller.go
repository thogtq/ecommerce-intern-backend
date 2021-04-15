package controllers

import (
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/constants"
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
	if len(productData.Images) == 0 {
		productData.Images = []string{"http://"+c.Request.Host + constants.PRODUCT_DEFAULT_IMAGE_PATH}
	} else {
		for index, image := range productData.Images {
			imageName := image[strings.LastIndex(image, "/")+1:]
			err := os.Rename(constants.PRODUCT_IMAGE_TEMP_DIR+imageName, constants.PRODUCT_IMAGE_DIR+imageName)
			if err != nil {
				c.Error(errors.ErrInternal(err.Error()))
				return
			}
			productData.Images[index] = strings.Replace(image, constants.PRODUCT_IMAGE_TEMP_URL, constants.PRODUCT_IMAGE_URL, 1)
		}

	}
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
	fileUrl := "http://" + c.Request.Host + constants.PRODUCT_IMAGE_TEMP_URL + fileName
	c.JSON(200, SuccessResponse(gin.H{"fileUrl": fileUrl}))

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
		c.Error(errors.ErrProductNotFound)
		return
	}
	products, err := productDAO.GetProductByID(c, objectID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, SuccessResponse(gin.H{"product": products}))
}
