package routes

import (
	"fmt"
	"net/http"

	"github.com/deveshpatel0101/jwt_jam/controllers"
	"github.com/deveshpatel0101/jwt_jam/keys"
	"github.com/julienschmidt/httprouter"
)

func GtUserAuth(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	tokenCookie, err := req.Cookie("token")
	username := p.ByName("username")
	var token string

	if err != nil {
		// generate new token
		token, err = controllers.GenToken(username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, err.Error())
			return
		}
	} else {
		token, err = controllers.Authenticate(tokenCookie.Value, "auth")
		if err != nil {
			// generate new token
			token, err = controllers.GenToken(username)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, err.Error())
				return
			}
		}
	}

	cookie := http.Cookie{Name: "token", Value: token, HttpOnly: true, Path: "/", Secure: true}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, string(keys.PubKey))
}
