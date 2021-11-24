package controllers

import (
	"time"

	"github.com/deveshpatel0101/jwt_jam/keys"
	"github.com/golang-jwt/jwt"
)

func GenToken(username string) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(keys.PrvKey)
	if err != nil {
		return "", err
	}

	claims := make(jwt.MapClaims)
	claims["username"] = username
	claims["auth"] = 1
	claims["verify"] = 0
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)

	if err != nil {
		return "", err
	}

	return token, nil
}
