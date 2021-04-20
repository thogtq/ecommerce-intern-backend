package dao

import (
	"context"

	"github.com/thogtq/ecommerce-server/database"
	"github.com/thogtq/ecommerce-server/errors"
	"github.com/thogtq/ecommerce-server/models"
	"github.com/thogtq/ecommerce-server/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDAO struct {
	userCollection *mongo.Collection
}

func (ud *UserDAO) New() *UserDAO {
	ud.userCollection = database.DBClient.Database("ecommerce").Collection("users")
	return ud
}
func (ud *UserDAO) Init() {
	ud.userCollection = database.DBClient.Database("ecommerce").Collection("users")
}
//Fix me
func (ud *UserDAO) CreateUser(ctx context.Context, userData *models.User) (insertID string, err error) {
	ud.Init()
	user := models.User{}
	if checkEmail := user.CheckIfEmailExist(ctx, ud.userCollection, userData.Email); checkEmail {
		return "", errors.ErrEmailExisted
	}
	userData.Password = user.HashPassword(userData.Password)
	userData.Role = "user"
	res, err := ud.userCollection.InsertOne(ctx, userData)
	if err != nil {
		return "", errors.ErrInternal(err.Error())
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
func (ud *UserDAO) Login(ctx context.Context, loginData *models.UserLogin) (user *models.User, err error) {
	ud.Init()
	user = &models.User{}
	result := ud.userCollection.FindOne(ctx, bson.M{"email": loginData.Email})
	err = result.Decode(user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.ErrEmailNotFound
	}
	if err != nil {
		return nil, errors.ErrInternal(err.Error())
	}
	if checkPassword := user.VerifyPassword(user.Password, loginData.Password); !checkPassword {
		return nil, errors.ErrInvalidPassword
	}
	user.Password = ""
	return user, nil
}

func (ud *UserDAO) GetUserByUserID(ctx context.Context, userID string) (user *models.User, err error) {
	ud.Init()
	user = &models.User{}
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.ErrInternal(err.Error())
	}
	result := ud.userCollection.FindOne(ctx, bson.M{"_id": objectID})
	if err := result.Decode(user); err != nil {
		errors.ErrInternal(err.Error())
	}
	return user, nil
}
func (ud *UserDAO) UpdateUser(c context.Context, userData *models.User, userID string) error {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.ErrInvalidUserID
	}
	filter := bson.D{{Key: "_id", Value: objectID}}
	fields := bson.D{}
	if userData.Fullname != "" {
		fields = append(fields, bson.E{"fullname", userData.Fullname})
	}
	if userData.Email != "" {
		fields = append(fields, bson.E{"email", userData.Email})
	}
	update := bson.D{
		{Key: "$set", Value: fields},
	}
	result, err := ud.userCollection.UpdateOne(c, filter, update)
	if err != nil {
		return errors.ErrInternal(err.Error())
	}
	if result.MatchedCount == 0 {
		return errors.ErrUserNotFound
	}
	return nil
}
func (ud *UserDAO) UpdateUserPassword(c context.Context, new, userID string) error {
	

	return nil
}
