package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vetusbs/gomino/models"
	"github.com/vetusbs/gomino/server"
	"github.com/vetusbs/gomino/views"
)

func game() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == http.MethodPost {
			data := views.CreateGameRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			game := models.InitGame(data)

			server.AddGame(&game)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			js, _ := json.Marshal(models.CreateGameDto(&game))
			w.Write(js)
			fmt.Println("END")
		} else if r.Method == http.MethodGet {
			fmt.Println("START")
			id := mux.Vars(r)["id"]

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			game := server.GetGame(id)
			js, _ := json.Marshal(models.CreateGameDto(game))
			w.Write(js)
			fmt.Println("END")
		} else if r.Method == http.MethodPut {
			fmt.Printf("Method put start")
			data := views.ActionRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			game := server.GetGame(data.Game)
			fmt.Printf("Method put for player %v %v", game.GetCurrentPlayer(), data)

			if data.Type == "play" {
				isLeft := bool(data.Details["isLeft"].(bool))

				if err := game.PlayCard(
					game.GetCurrentPlayer(),
					models.CardDto{
						Left:  int(data.Details["left"].(float64)),
						Right: int(data.Details["right"].(float64)),
					}, isLeft); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(err.Error()))
					return
				}
			} else if data.Type == "pick" {
				fmt.Printf("current player %v", game.GetCurrentPlayer())
				game.Pick(game.GetCurrentPlayer())
			} else if data.Type == "addUser" {
				userID := string(data.Details["userId"].(string))
				userName := string(data.Details["userName"].(string))
				if err := game.AddUser(userID, userName); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(err.Error()))
					return
				}
			} else {

			}
			w.WriteHeader(http.StatusAccepted)
			js, _ := json.Marshal(models.CreateGameDto(game))
			w.Write(js)
			fmt.Println("END")
		}
	}
}
