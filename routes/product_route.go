package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/controllers"
)

func ProductPublicRoute(r *gin.RouterGroup) {
	//Static == GET
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/", controllers.GetProduct)
	r.Static("/products/image", "./files/images/products")
	r.Static("/products/temp/", "./files/images/temp")
	//Legacy DB support
	r.Static("/product/image", "./files/images/products")
	r.Static("/product/temp/", "./files/images/temp")
}
func ProductAdminRoute(r *gin.RouterGroup) {
	r.DELETE("/products/", controllers.DeleteProduct)

	r.POST("/products/image", controllers.UploadProductImage)
	r.POST("/products/", controllers.CreateProduct)
	r.PUT("/products/", controllers.UpdateProduct)
}
