package helpers

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/thogtq/ecommerce-server/errors"
)

type SignedDetails struct {
	UserID string
	Role   string
	jwt.StandardClaims
}

func GenerateTokens(userID, role string) (string, string, int64, error) {
	tokenExpireAt := time.Now().Local().Add(time.Hour * time.Duration(24)).Unix()
	refreshTokenExpireAt := time.Now().Local().Add(time.Hour * time.Duration(168)).Unix()
	SECRET_KEY := os.Getenv("JWT_SECRET")
	if SECRET_KEY == "" {
		log.Panicf("unable to load jwt secret key")
	}
	claims := &SignedDetails{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpireAt,
		},
	}
	refreshClaims := &SignedDetails{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshTokenExpireAt,
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", 0, errors.ErrInternal(err.Error())
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", 0, errors.ErrInternal(err.Error())
	}
	return token, refreshToken, tokenExpireAt, nil
}
func ValidateToken(signedToken string) (*SignedDetails, error) {
	SECRET_KEY := os.Getenv("JWT_SECRET")
	if SECRET_KEY == "" {
		log.Panicf("unable to load jwt secret key")
	}
	token, err := jwt.ParseWithClaims(
		signedToken, &SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		if err.(*jwt.ValidationError).Errors == jwt.ValidationErrorExpired {
			return nil, errors.ErrExpiredToken
		}
		return nil, errors.ErrInvalidToken
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return nil, errors.ErrInvalidToken
	}
	return claims, nil
}
