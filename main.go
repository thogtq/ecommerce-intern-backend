package main

import (
	// "context"
	// "fmt"

	"fmt"

	"github.com/subosito/gotenv"
	"github.com/thogtq/ecommerce-server/database"
	"github.com/thogtq/ecommerce-server/models"
)

func init() {
	gotenv.Load()
}
func main() {
	database.Connect()
	defer database.Disconnect()
	
	fmt.Println(models.Test())
}
