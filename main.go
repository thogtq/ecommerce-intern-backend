package main

import (
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	"github.com/thogtq/ecommerce-server/database"
	"github.com/thogtq/ecommerce-server/middlewares"
	"github.com/thogtq/ecommerce-server/routes"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	gotenv.Load()
}
func main() {
	database.Connect()
	defer database.Disconnect()
	r := gin.Default()
	
	r.GET("/test", func(c *gin.Context) {
		b := bson.D{bson.E{"_id", "12313"}}
		e := bson.E{"name", "hahaha"}
		b = append(b, e)

		c.JSON(200, b)
	})
	//
	//Middlewares
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.ErrorHandler())
	r.Use(middlewares.ContextHandler())
	//Group
	public := r.Group("/api")
	public.Use()
	{
		routes.UserPublicRoute(public)
		routes.ProductPublicRoute(public)
	}
	authorized := r.Group("/api")
	authorized.Use(middlewares.AuthRequired())
	{
		routes.UserPrivateRoute(authorized)
	}
	admin := r.Group("/api")
	admin.Use(middlewares.AdminAuthRequired())
	{
		routes.UserAdminRoute(admin)
		routes.ProductAdminRoute(admin)
	}
	r.Run(":8080")
}
