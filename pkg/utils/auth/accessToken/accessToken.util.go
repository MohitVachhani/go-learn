package accesstokenutil

import (
	"fmt"
	"log"
	"time"

	envUtil "github.com/MohitVachhani/go-learn/pkg/utils/env"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(emailId string) string {
	var jwtSecret = envUtil.Get("JWT_SECRET")

	atClaims := jwt.MapClaims{}

	atClaims["emailId"] = emailId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	atClaims["tokenUse"] = "accessToken"

	at := jwt.NewWithClaims(jwt.SigningMethodES256, atClaims)

	token, err := at.SignedString([]byte(jwtSecret))

	if err != nil {
		log.Fatal("Error while creating token", err)
	}

	fmt.Println(token)
	return token
}
