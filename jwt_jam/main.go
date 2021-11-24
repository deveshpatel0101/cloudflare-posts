package main

import (
	"fmt"
	"net/http"

	"github.com/deveshpatel0101/jwt_jam/routes"
	"github.com/julienschmidt/httprouter"
)

func main() {
	mux := httprouter.New()
	mux.GET("/auth/:username", routes.GtUserAuth)
	mux.GET("/verify/", routes.GtVerify)
	mux.GET("/stats/", routes.Stats)
	mux.GET("/README.txt/", routes.ReadMe)
	fmt.Println("Server started on port 8000...")
	http.ListenAndServe(":8000", mux)
}
