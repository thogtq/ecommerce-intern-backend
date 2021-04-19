package controllers

import (
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/dao"
	"github.com/thogtq/ecommerce-server/models"
	"github.com/thogtq/ecommerce-server/pkg/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var orderDAO dao.OrderDAO

func CreateOrder(c *gin.Context) {
	var (
		err      error
		order    = &models.Order{}
		subtotal int
	)
	//Bind and set default value
	c.BindJSON(order)
	order.OrderDate = time.Now()
	order.OrderID = uuid.NewShortID()
	order.UserID, err = GetContextUserID(c)
	order.Status = "Pending"
	if err != nil {
		c.Error(err)
		return
	}
	//Calculate product amount(price*quantity) and subtotal(sum of amount)
	for index, product := range order.ProductsOrder {
		objectID, _ := primitive.ObjectIDFromHex(product.ProductID)
		_product, err := productDAO.GetProductByID(c.Request.Context(), objectID)
		if err != nil {
			c.Error(err)
			return
		}
		amount := product.Quantity * _product.Price
		order.ProductsOrder[index].Amount = amount
		subtotal += amount
	}
	order.Subtotal = subtotal
	//Performs insert db
	id, err := orderDAO.InsertOrder(c.Request.Context(), order)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, SuccessResponse(gin.H{"_id": id, "orderID": order.OrderID}))
}
func GetOrders(c *gin.Context) {
	var (
		filter = &models.OrderFilter{}
	)
	c.Bind(filter)
	if filter.Page == 0 {
		filter.Page = 1
	}
	if filter.Limit == 0 {
		filter.Limit = -1
	}
	orders, counts, err := orderDAO.GetOrders(c.Request.Context(), filter)
	if err != nil {
		c.Error(err)
		return
	}
	pages := math.Ceil(float64(counts) / float64(filter.Limit))
	c.JSON(200, SuccessResponse(
		gin.H{
			"orders":  orders,
			"counts": counts,
			"pages":  pages,
		},
	))
}