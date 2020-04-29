package controller

import (
	"github.com/gorilla/mux"
)

func Register() *mux.Router {
	api := mux.NewRouter()
	api.HandleFunc("/game/{id}", game())
	api.HandleFunc("/game", game())
	return api
}
