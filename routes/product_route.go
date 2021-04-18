package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/controllers"
)

func ProductPublicRoute(r *gin.RouterGroup) {
	//Static == GET
	r.GET("/products/", controllers.GetProducts)
	r.GET("/product/", controllers.GetProduct)
	r.Static("/product/image", "./files/images/products")
	r.Static("/product/temp/", "./files/images/temp")
}
func ProductAdminRoute(r *gin.RouterGroup) {
	r.DELETE("/product/", controllers.DeleteProduct)

	r.POST("/product/image", controllers.UploadProductImage)
	r.POST("/product", controllers.CreateProduct)
	r.PUT("/product", controllers.UpdateProduct)
}
