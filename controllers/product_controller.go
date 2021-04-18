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
		productData.Images = []string{"http://" + c.Request.Host + constants.PRODUCT_DEFAULT_IMAGE_PATH}
	} else {
		for index, image := range productData.Images {
			imageName := services.GetProductImageNameFromURL(image)
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
	filter := &models.ProductFilters{}
	err := c.Bind(filter)
	if err != nil {
		c.Error(errors.ErrInternal(err.Error()))
	}
	if filter.SortOrder == 0 {
		filter.SortOrder = -1
	}
	products, err := productDAO.GetProducts(c.Request.Context(), filter)
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
	products, err := productDAO.GetProductByID(c.Request.Context(), objectID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, SuccessResponse(gin.H{"product": products}))
}
func DeleteProduct(c *gin.Context) {
	productID := c.Request.URL.Query().Get("productID")
	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		c.Error(errors.ErrProductNotFound)
		return
	}
	product, err := productDAO.GetProductByID(c.Request.Context(), objectID)
	if err != nil {
		c.Error(err)
		return
	}
	err = productDAO.DeleteProduct(c.Request.Context(), objectID)
	if err != nil {
		c.Error(err)
		return
	}
	if !strings.Contains(product.Images[0], constants.PRODUCT_DEFAULT_IMAGE_PATH) {
		for _, image := range product.Images {
			imageName := services.GetProductImageNameFromURL(image)
			err = os.Remove(constants.PRODUCT_IMAGE_DIR + imageName)
		}
	}
	c.JSON(200, SuccessResponse(gin.H{"result": "deleted"}))
}
func UpdateProduct(c *gin.Context) {
	productData := &models.Product{}
	c.BindJSON(productData)
	productData.ParentCategories = services.GetParentCategories(productData.Categories)
	//Set default image if no product images found
	if len(productData.Images) == 0 {
		productData.Images = []string{"http://" + c.Request.Host + constants.PRODUCT_DEFAULT_IMAGE_PATH}
	} else {
		for index, image := range productData.Images {
			//Keep existing image
			if strings.Contains(image, constants.PRODUCT_IMAGE_URL) {
				continue
			}
			imageName := services.GetProductImageNameFromURL(image)
			err := os.Rename(constants.PRODUCT_IMAGE_TEMP_DIR+imageName, constants.PRODUCT_IMAGE_DIR+imageName)
			if err != nil {
				c.Error(errors.ErrInternal(err.Error()))
				return
			}
			productData.Images[index] = strings.Replace(image, constants.PRODUCT_IMAGE_TEMP_URL, constants.PRODUCT_IMAGE_URL, 1)
		}

	}
	err := productDAO.UpdateProduct(c.Request.Context(), productData)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, SuccessResponse(gin.H{"result": "updated"}))
}
