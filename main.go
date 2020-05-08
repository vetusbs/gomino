package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/vetusbs/gomino/controller"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	mux := controller.Register()

	fmt.Printf("listening in port %v \n", port)
	//http.ListenAndServeTLS(":"+port, "certs/server.crt", "certs/server.key", mux)
	http.ListenAndServe(":"+port, mux)
}
