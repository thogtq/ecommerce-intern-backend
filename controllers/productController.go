package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/models"
)

var productModel models.Product

func AddProduct(c *gin.Context) {
	productData := &models.Product{}
	c.BindJSON(productData)
	res, err := productModel.InsertProduct(productData)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status":    "posted",
		"message":   "product added",
		"productID": res,
	})
}
