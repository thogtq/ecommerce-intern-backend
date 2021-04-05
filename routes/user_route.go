package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/controllers"
)

func UserRoute(r *gin.RouterGroup) {
	r.POST("/user/login", controllers.Login)
	r.POST("/user", controllers.Regiser)
}
func UserAuthorizedRoute(r *gin.RouterGroup) {
	r.PUT("/user", controllers.UpdateUser)
}
