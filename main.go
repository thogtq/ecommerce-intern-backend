package main

import (
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	"github.com/thogtq/ecommerce-server/database"
	"github.com/thogtq/ecommerce-server/middlewares"
	"github.com/thogtq/ecommerce-server/routes"
)

func init() {
	gotenv.Load()
}
func main() {
	database.Connect()
	defer database.Disconnect()
	r := gin.Default()
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
		routes.ProductPrivateRoute(authorized)
	}
	admin := r.Group("/api")
	admin.Use(middlewares.AdminAuthRequired())
	{
		routes.UserAdminRoute(admin)
	}
	r.Run(":8080")
}
