package dao

import (
	"context"

	"github.com/thogtq/ecommerce-server/database"
	"github.com/thogtq/ecommerce-server/errors"
	"github.com/thogtq/ecommerce-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReviewDAO struct {
	reviewCollection *mongo.Collection
}

func (rd *ReviewDAO) New() *ReviewDAO {
	rd.reviewCollection = database.DBClient.Database("ecommerce").Collection("reviews")
	return rd
}

func (rd *ReviewDAO) InsertReview(c context.Context, review *models.Review) (string, error) {
	result, err := rd.reviewCollection.InsertOne(c, review)
	if err != nil {
		return "", errors.ErrInternal(err.Error())
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
func (rd *ReviewDAO) DeleteReview(c context.Context, reviewID string) error {
	objectID, err := primitive.ObjectIDFromHex(reviewID)
	if err != nil {
		return errors.ErrReviewNotFound
	}
	result, err := rd.reviewCollection.DeleteOne(c, bson.M{"_id": objectID})
	if err != nil {
		return errors.ErrInternal(err.Error())
	}
	if result.DeletedCount == 0 {
		return errors.ErrReviewNotFound
	}
	return nil
}
func (rd *ReviewDAO) GetReviewsByProductID(c context.Context, productID string, filter *models.ReviewFilters) (*[]models.Review, int64, error) {
	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, 0, errors.ErrProductNotFound
	}
	reviews := []models.Review{}
	skipRows := (filter.Page - 1) * filter.Limit
	findFilter := bson.M{"productID": objectID}
	findOptions := options.Find()
	findOptions.SetSort(bson.M{"reviewDate": -1})
	findOptions.SetLimit(int64(filter.Limit))
	findOptions.SetSkip(int64(skipRows))
	cur, err := rd.reviewCollection.Find(c, findFilter, findOptions)
	if err != nil {
		return nil, 0, errors.ErrInternal(err.Error())
	}
	counts, _ := rd.reviewCollection.CountDocuments(c, findFilter)
	for cur.Next(c) {
		review := &models.Review{}
		err := cur.Decode(review)
		if err != nil {
			return nil, 0, errors.ErrInternal(err.Error())
		}
		reviews = append(reviews, *review)
	}
	return &reviews, counts, nil
}
func (rd *ReviewDAO) GetReviewByID(c context.Context, reviewID string) (*models.Review, error) {
	review := &models.Review{}
	objectID, err := primitive.ObjectIDFromHex(reviewID)
	if err != nil {
		return nil, errors.ErrReviewNotFound
	}
	findFilter := bson.M{"_id": objectID}
	result := rd.reviewCollection.FindOne(c, findFilter)
	if result.Err() == mongo.ErrNoDocuments {
		return nil, errors.ErrReviewNotFound
	}
	err = result.Decode(review)
	if err != nil {
		return nil, errors.ErrInternal(err.Error())
	}
	return review, nil
}
