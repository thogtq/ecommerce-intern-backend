package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/controllers"
)

func UserPublicRoute(r *gin.RouterGroup) {
	r.POST("/user/login", controllers.Login)
	//Gop voi user/login
	r.POST("/admin/login", controllers.AdminLogin)
	r.POST("/users/", controllers.Regiser)
	r.Static("/user/image/", "./files/images/users")
}
func UserPrivateRoute(r *gin.RouterGroup) {
	r.PUT("/users/password", controllers.UpdateUserPassword)
	r.GET("/users/", controllers.GetUser)
	r.GET("/users/token", controllers.GetNewAccessToken)
	r.PUT("/users/", controllers.UpdateUser)
}
func UserAdminRoute(r *gin.RouterGroup) {
	// r.GET("/users", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"Get many users": 1})
	// })
}
