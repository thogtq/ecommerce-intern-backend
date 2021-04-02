package dao

import (
	"context"

	"github.com/thogtq/ecommerce-server/database"
	"github.com/thogtq/ecommerce-server/errors"
	"github.com/thogtq/ecommerce-server/helpers"
	"github.com/thogtq/ecommerce-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDAO struct {
	UserCollection *mongo.Collection
}

func (ud *UserDAO) Init() {
	ud.UserCollection = database.DBClient.Database("ecommerce").Collection("users")
}

func (ud *UserDAO) Register(ctx context.Context, userData *models.User) (insertID string, err error) {
	ud.Init()
	user := models.User{}

	if checkEmail := user.CheckIfEmailExist(ctx, ud.UserCollection, userData.Email); checkEmail {
		return "", errors.ErrEmailExisted
	}
	userData.Password = user.HashPassword(userData.Password)
	res, err := ud.UserCollection.InsertOne(ctx, userData)
	if err != nil {
		return "", errors.ErrNotInserted
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
func (ud *UserDAO) Login(ctx context.Context, loginData *models.UserLogin) (user *models.User, token *models.UserToken, err error) {
	ud.Init()
	token = &models.UserToken{}
	user = &models.User{}

	result := ud.UserCollection.FindOne(ctx, bson.M{"email": loginData.Email})
	if err := result.Decode(user); err == mongo.ErrNoDocuments {
		return nil, nil, errors.ErrEmailNotExisted
	}
	if checkPassword := user.VerifyPassword(user.Password, loginData.Password); !checkPassword {
		return nil, nil, errors.ErrInvalidPassword
	}
	token.AccessToken, token.RefreshToken, err = helpers.GenerateTokens(user.UserID.Hex(), user.Email)
	if err != nil {
		return nil, nil, errors.ErrTokenNotGenerated
	}
	user.Password = ""
	return user, token, nil
}
