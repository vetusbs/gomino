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
	panic(http.ListenAndServe(":"+port, mux))
}
