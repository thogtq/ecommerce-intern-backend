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
	//Middlewares
	r.Use(middlewares.CORSMiddleware())
	//Group
	public := r.Group("/api")
	public.Use()
	{
		routes.UserRoute(public)
		public.Static("/product/image", "./files/images/products")
		//routes.ProductRoute(public)
	}
	authorized := r.Group("/api")
	authorized.Use(middlewares.AuthRequired())
	{
		routes.UserAuthorizedRoute(authorized)
	}
	r.Run(":8080")
}
