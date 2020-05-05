package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/vetusbs/gomino/views"

	"github.com/vetusbs/gomino/controller"
	"github.com/vetusbs/gomino/models"
	"github.com/vetusbs/gomino/server"
)

var hub server.Hub

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	mux := controller.Register()
	hub := server.InitHub()
	// should be a post
	type test struct {
		Id string `json`
	}

	mux.HandleFunc("/createGame", func(w http.ResponseWriter, request *http.Request) {
		game := models.InitGame(views.CreateGameRequest{Players: 3})
		hub.Games[game.GetId()] = &game

		js, _ := json.Marshal(models.CreateGameDto(&game))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	})

	// get Game
	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {

		name := r.URL.Query().Get("name")

		js, _ := json.Marshal(models.CreateGameDto(hub.Games[name]))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	})

	http.ListenAndServeTLS(":"+port, "certs/server.crt", "certs/server.key", mux)
}

func main0() {
	game := models.InitGame(views.CreateGameRequest{Players: 3})

	game.PrintGameState()
	var currentPlayer *models.Player
	for game.IsFinished() == false {
		currentPlayer = game.GetCurrentPlayer()
		var i int
		var head bool
		fmt.Printf("%v: ", currentPlayer.GetName())
		fmt.Scan(&i)

		if i == 33 {
			currentPlayer.AutoPlay(&game)
		} else {
			fmt.Scan(&head)
			fmt.Printf("%v %v %v\n", currentPlayer.GetName(), i, head)
			if err := game.PlayCardPublic(currentPlayer, i, head); err != nil {
				fmt.Println(err)
				game.Pick(currentPlayer)
			}
		}

		if i < 0 {
			break
		}

		game.PrintGameState()
		fmt.Println("*******************************")
	}
}
