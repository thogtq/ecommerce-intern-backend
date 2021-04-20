package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/controllers"
)

func OrderPrivateRoute(r *gin.RouterGroup) {
	r.POST("/orders", controllers.CreateOrder)
}
func OrderAdminRoute(r *gin.RouterGroup) {
	r.PUT("/orders/status",controllers.UpdateOrderStatus)
	r.GET("/orders", controllers.GetOrders)
}
