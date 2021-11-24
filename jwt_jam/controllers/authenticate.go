package controllers

import (
	"errors"
	"fmt"

	"github.com/deveshpatel0101/jwt_jam/keys"
	"github.com/golang-jwt/jwt"
)

// increments "auth" or "verify" by one if token is valid and returns the update token
func Authenticate(token, route string) (string, error) {
	// parse the public key
	rsaPubKey, err := jwt.ParseRSAPublicKeyFromPEM(keys.PubKey)
	if err != nil {
		return "", err
	}

	// validate the token
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Invalid token")
		}
		return rsaPubKey, nil
	})

	if err != nil {
		return "", errors.New("Invalid token!")
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return "", errors.New("Invalid token!")
	}

	// increment the stats
	claims[route] = claims[route].(float64) + 1

	// parse the private key
	rsaPrvKey, err := jwt.ParseRSAPrivateKeyFromPEM(keys.PrvKey)
	if err != nil {
		return "", err
	}

	// re-generate the token with updated stats
	token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(rsaPrvKey)
	if err != nil {
		return "", err
	}

	// return the token
	return token, nil
}
