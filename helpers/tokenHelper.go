package helpers

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	UserID string
	Email  string
	jwt.StandardClaims
}

func GenerateTokens(userID, email string) (string, string, error) {
	SECRET_KEY := os.Getenv("JWT_SECRET")
	if SECRET_KEY == "" {
		return "", "", fmt.Errorf("unable to load jwt secret key")
	}
	claims := &SignedDetails{
		UserID: userID,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}
	fmt.Println(token)
	fmt.Println(SECRET_KEY)
	return token, refreshToken, nil
}
func ValidateToken(signedToken string) (*SignedDetails, error) {
	SECRET_KEY := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(
		signedToken, &SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, fmt.Errorf("expired token")
	}
	return claims, nil
}

// func UpdateTokens(signedToken, refreshToken, userID string, userCol *mongo.Collection) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
// 	defer cancel()
// 	var updateObj primitive.D
// 	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken})
// 	updateObj = append(updateObj, bson.E{Key: "refreshToken", Value: refreshToken})
// 	upsert := true
// 	filter := bson.M{"_id": userID}
// 	opt := &options.UpdateOptions{
// 		Upsert: &upsert,
// 	}
// 	_, err := userCol.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}}, opt)
// 	if err != nil {
// 		return fmt.Errorf("can not update token")
// 	}
// 	return nil
// }
