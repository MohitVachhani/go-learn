package accesstokenutil

import (
	"log"
	"time"

	authInterface "github.com/MohitVachhani/go-learn/pkg/structs/auth"
	envUtil "github.com/MohitVachhani/go-learn/pkg/utils/env"

	"github.com/dgrijalva/jwt-go"
)

func CreateAccessToken(payload authInterface.AccessTokenPayload) string {
	var jwtSecret = envUtil.Get("JWT_SECRET")

	atClaims := jwt.MapClaims{}

	atClaims["emailId"] = payload.EmailId
	atClaims["userId"] = payload.UserId
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	atClaims["tokenUse"] = payload.TokenUse

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(jwtSecret))

	if err != nil {
		log.Fatal("Error while creating access token", err)
	}

	return token
}

func CreateRefreshToken(payload authInterface.RefreshTokenPayload) string {
	var jwtSecret = envUtil.Get("JWT_SECRET")

	atClaims := jwt.MapClaims{}

	atClaims["userId"] = payload.UserId
	atClaims["exp"] = time.Now().Add(time.Hour * 24 * 3).Unix()
	atClaims["tokenUse"] = payload.TokenUse

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(jwtSecret))

	if err != nil {
		log.Fatal("Error while creating refresh token", err)
	}

	return token
}

func VerifyAccessToken(accessToken string) (authInterface.AccessTokenPayload, string) {

	var jwtSecret = envUtil.Get("JWT_SECRET")

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return authInterface.AccessTokenPayload{}, "INVALID_TOKEN"
	}

	if token.Valid == false {
		return authInterface.AccessTokenPayload{}, "NOT_VALID"
	}

	if claims["tokenUse"] != "accessToken" {
		return authInterface.AccessTokenPayload{}, "NOT_ACCESS_TOKEN"
	}

	var emailId = claims["emailId"].(string)
	var userId = claims["userId"].(string)

	return authInterface.AccessTokenPayload{
		EmailId: emailId,
		UserId:  userId,
	}, ""
}
