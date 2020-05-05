package controller

import (
	"github.com/gorilla/mux"
)

// Register is the main function of the http server
func Register() *mux.Router {
	api := mux.NewRouter()
	api.HandleFunc("/game/{id}", game())
	api.HandleFunc("/game", game())
	return api
}
