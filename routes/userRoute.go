package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/controllers"
)

func UserRoute(r *gin.Engine){
	r.POST("api/user/login", controllers.Login)
	r.POST("api/user", controllers.Regiser)
}