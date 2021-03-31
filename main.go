package main

import (
	// "context"
	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	"github.com/thogtq/ecommerce-server/controllers"
	"github.com/thogtq/ecommerce-server/database"
)

func init() {
	gotenv.Load()
}
func main() {
	database.Connect()
	defer database.Disconnect()
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.POST("api/user/login", controllers.Login)
	r.POST("api/user", controllers.Regiser)
	r.POST("api/product", controllers.AddProduct)
	r.Run(":8080")
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
