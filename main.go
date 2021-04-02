package main

import (
	// "context"
	// "fmt"

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
	//Routes
	routes.UserRoute(r)
	routes.ProductRoute(r)
	r.Static("/files/images", "./files/images/")
	r.Run(":8080")
}
