package controllers

import (
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/dao"
	"github.com/thogtq/ecommerce-server/errors"
	"github.com/thogtq/ecommerce-server/models"
	"github.com/thogtq/ecommerce-server/pkg/uuid"
	"github.com/thogtq/ecommerce-server/services"
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
	if len(order.ProductsOrder) == 0 {
		c.Error(errors.ErrEmptyCart)
		return
	}
	//Calculate product amount(price*quantity) and subtotal(sum of amount)
	for index, product := range order.ProductsOrder {
		objectID, _ := primitive.ObjectIDFromHex(product.ProductID)
		_product, err := productDAO.New().GetProductByID(c.Request.Context(), objectID)
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
	id, err := orderDAO.New().InsertOrder(c.Request.Context(), order)
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
	orders, counts, err := orderDAO.New().GetOrders(c.Request.Context(), filter)
	if err != nil {
		c.Error(err)
		return
	}
	pages := math.Ceil(float64(counts) / float64(filter.Limit))
	c.JSON(200, SuccessResponse(
		gin.H{
			"orders": orders,
			"counts": counts,
			"pages":  pages,
		},
	))
}
func UpdateOrderStatus(c *gin.Context) {
	var (
		orderID    = c.Request.URL.Query().Get("orderID")
		bodyParams struct {
			Status string `json:"status"`
		}
		statusDB = []string{"Pending", "Completed", "Canceled"}
	)
	if orderID == "" {
		c.Error(errors.ErrNoOrderID)
		return
	}
	err := c.BindJSON(&bodyParams)
	if err != nil {
		c.Error(errors.ErrInternal(err.Error()))
		return
	}
	if !services.Contains(statusDB, bodyParams.Status) {
		c.Error(errors.ErrInvalidOrderStatus)
		return
	}
	order, err := orderDAO.New().GetOrderByID(c.Request.Context(), orderID)
	if err != nil {
		c.Error(err)
		return
	}
	if order.Status != "Pending" || bodyParams.Status != "Canceled" {
		for _, product := range order.ProductsOrder {
			sold := product.Quantity
			profit := product.Amount
			if order.Status == "Completed" && bodyParams.Status == "Canceled" {
				sold *= -1
				profit *= -1
			}
			err := productDAO.UpdateProductSale(c.Request.Context(), product.ProductID, sold, profit)
			if err != nil {
				c.Error(err)
				return
			}
		}
	}
	err = orderDAO.New().UpdateStatus(c.Request.Context(), orderID, bodyParams.Status)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, SuccessResponse(gin.H{"result": "updated"}))
}
