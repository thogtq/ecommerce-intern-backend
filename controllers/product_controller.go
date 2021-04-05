package controllers

import (
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/dao"
	"github.com/thogtq/ecommerce-server/models"
)

var productDAO dao.ProductDAO

func CreateProduct(c *gin.Context) {
	productData := &models.Product{}
	c.BindJSON(productData)
	productData.CreatedAt = time.Now()
	res, err := productDAO.InsertProduct(c, productData)
	if err != nil {
		//error
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	//Fix me
	c.JSON(200, gin.H{
		"status":    "posted",
		"message":   "product added",
		"productID": res,
	})
}
func UploadProductImage(c *gin.Context) {
	file, err := c.FormFile("productImage")
	file.Filename = strings.ToLower(file.Filename)
	fileExt := path.Ext(file.Filename)
	acceptedExt := []string{".jpg", ".png", ".jpeg", ".gif"}
	if !contains(acceptedExt, fileExt) {
		//error
		c.JSON(400, "Invalid file type")
		return
	}
	if err != nil {
		//error
		c.JSON(400, err.Error())
		return
	}
	fileName, err := productDAO.UploadImage(c, file)
	if err != nil {
		//error
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"status":   "success",
		"fileName": fileName,
	})
}
