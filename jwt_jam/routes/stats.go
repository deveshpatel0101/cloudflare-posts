package routes

import (
	"fmt"
	"net/http"

	"github.com/deveshpatel0101/jwt_jam/controllers"
	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
)

func Stats(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	tokenCookie, err := req.Cookie("token")

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Token not found")
		return
	}

	_, err = controllers.Authenticate(tokenCookie.Value, "auth")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid token")
		return
	}

	jwtToken, _ := jwt.Parse(tokenCookie.Value, nil)
	claims, _ := jwtToken.Claims.(jwt.MapClaims)
	authVisits := claims["auth"].(float64)
	verifyVisits := claims["verify"].(float64)
	fmt.Fprintf(w, "Auth Visits: %v\nVerify Visits: %v", authVisits, verifyVisits)
}
