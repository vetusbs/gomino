package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Register() *mux.Router {
	api := mux.NewRouter()
	api.HandleFunc("/game/{id}", game()).Methods(http.MethodGet)
	api.HandleFunc("/game", game())
	return api
}
