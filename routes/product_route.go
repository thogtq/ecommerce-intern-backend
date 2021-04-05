package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/controllers"
)

func ProductPublicRoute(r *gin.RouterGroup) {
	//Static == GET
	r.Static("/product/image", "./files/images/products")
}
func ProductPrivateRoute(r *gin.RouterGroup) {
	r.POST("/product/image", controllers.UploadProductImage)
	r.POST("/product", controllers.CreateProduct)
}
