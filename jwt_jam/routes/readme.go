package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ReadMe(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	content, err := ioutil.ReadFile("README.txt")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid token")
		return
	}

	fmt.Fprintf(w, string(content))
	return
}
