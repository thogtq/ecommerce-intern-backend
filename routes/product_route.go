package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/controllers"
)

func ProductPublicRoute(r *gin.RouterGroup) {
	//Static == GET
	r.GET("/products/", controllers.GetProducts)
	r.Static("/product/image", "./files/images/products")
	r.Static("/product/temp/", "./files/images/temp")
}
func ProductAdminRoute(r *gin.RouterGroup) {
	r.POST("/product/image", controllers.UploadProductImage)
	r.POST("/product", controllers.CreateProduct)
}
