package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/controllers"
)

func ReviewPublicRoute(r *gin.RouterGroup) {
	r.GET("/reviews", controllers.GetReviews)
}
func ReviewPrivateRoute(r *gin.RouterGroup) {
	r.DELETE("/reviews/", controllers.DeleteReview)
	r.POST("/reviews/", controllers.CreateReview)
}
