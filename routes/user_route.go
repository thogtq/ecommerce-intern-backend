package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/controllers"
)

func UserPublicRoute(r *gin.RouterGroup) {
	r.POST("/user/login", controllers.Login)
	r.POST("/admin/login", controllers.AdminLogin)
	r.POST("/user", controllers.Regiser)
}
func UserPrivateRoute(r *gin.RouterGroup) {
	r.PUT("/user", controllers.UpdateUser)
}
func UserAdminRoute(r *gin.RouterGroup) {

}
