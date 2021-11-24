package routes

import (
	"fmt"
	"net/http"

	"github.com/deveshpatel0101/jwt_jam/controllers"
	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
)

func GtVerify(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	tokenCookie, err := req.Cookie("token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Token not found")
		return
	}

	token, err := controllers.Authenticate(tokenCookie.Value, "verify")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid token")
		return
	}

	jwtToken, _ := jwt.Parse(token, nil)
	claims, _ := jwtToken.Claims.(jwt.MapClaims)

	cookie := http.Cookie{Name: "token", Value: token, HttpOnly: true, Path: "/", Secure: true}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, claims["username"].(string))
}
