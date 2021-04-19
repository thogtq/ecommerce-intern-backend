package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ProductID        primitive.ObjectID `bson:"_id,omitempty" json:"productID"`
	Name             string             `bson:"name" json:"name"`
	CreatedAt        time.Time          `bson:"createdAt" json:"createdAt"`
	Categories       []string           `bson:"categories" json:"categories"`
	ParentCategories []string           `bson:"parentCategories" json:"parentCategories"`
	Images           []string           `bson:"images" json:"images"`
	Brand            string             `bson:"brand" json:"brand"`
	Price            int                `bson:"price" json:"price"`
	Sizes            []string           `bson:"sizes" json:"sizes"`
	Colors           []string           `bson:"colors" json:"colors"`
	Quantity         int                `bson:"quantity" json:"quantity"`
	Sold             int                `bson:"sold" json:"sold"`
	Description      string             `bson:"description" json:"description"`
}
type ProductFilters struct {
	Limit     int    `form:"limit"`
	SortBy    string `form:"sortBy"`
	SortOrder int    `form:"sortOrder"`
	Search    string `form:"search"`
	Category  string `form:"category"`
	Size      string `form:"size"`
	Color     string `form:"color"`
	Brand     string `form:"brand"`
	MinPrice  int    `form:"minPrice"`
	MaxPrice  int    `form:"maxPrice"`
	Available string `form:"available"`
	Page      int    `form:"page"`
}
