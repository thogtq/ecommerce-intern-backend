package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/controllers"
)

func ProductRoute(r *gin.Engine) {
	r.POST("api/product", controllers.AddProduct)
}
