package controllers

import (
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/dao"
	"github.com/thogtq/ecommerce-server/errors"
	"github.com/thogtq/ecommerce-server/models"
)

var reviewDAO dao.ReviewDAO

func CreateReview(c *gin.Context) {
	review := &models.Review{}
	userID, err := GetContextUserID(c)
	if err != nil {
		c.Error(errors.ErrUserNotFound)
		return
	}
	user, err := userDAO.New().GetUserByUserID(c.Request.Context(), userID)
	if err != nil {
		c.Error(err)
		return
	}
	err = c.BindJSON(review)
	if err != nil {
		c.Error(errors.ErrInternal(err.Error()))
	}
	review.Fullname = user.Fullname
	review.ReviewDate = time.Now()
	review.UserID = userID
	err = productDAO.New().UpdateProductReview(c.Request.Context(), review.ProductID, review.Star, 1)
	if err != nil {
		c.Error(err)
		return
	}
	reviewID, err := reviewDAO.New().InsertReview(c.Request.Context(), review)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, SuccessResponse(gin.H{"reviewID": reviewID}))
}
func GetReviews(c *gin.Context) {
	productID := c.Request.URL.Query().Get("productID")
	if productID == "" {
		c.Error(errors.ErrInvalidParameters)
		return
	}
	filter := &models.ReviewFilters{}
	err := c.BindQuery(filter)
	if err != nil {
		c.Error(errors.ErrInvalidParameters)
		return
	}
	reviews, counts, err := reviewDAO.New().GetReviewsByProductID(c.Request.Context(), productID, filter)
	if err != nil {
		c.Error(err)
		return
	}
	pages := math.Ceil(float64(counts) / float64(filter.Limit))
	c.JSON(200, SuccessResponse(gin.H{"reviews": reviews, "pages": pages, "counts": counts}))
}
func DeleteReview(c *gin.Context) {
	userID, _ := GetContextUserID(c)
	reviewID := c.Request.URL.Query().Get("reviewID")
	if reviewID == "" {
		c.Error(errors.ErrReviewNotFound)
		return
	}
	review, err := reviewDAO.New().GetReviewByID(c.Request.Context(), reviewID)
	if err != nil {
		c.Error(err)
		return
	}
	if review.UserID != userID {
		c.Error(errors.ErrUnauthorized)
		return
	}
	review.Star *= -1
	err = productDAO.New().UpdateProductReview(c.Request.Context(), review.ProductID, review.Star, -1)
	if err != nil {
		c.Error(err)
		return
	}
	err = reviewDAO.New().DeleteReview(c.Request.Context(), reviewID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, SuccessResponse(gin.H{"result": "deleted"}))
}
