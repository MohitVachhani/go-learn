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
