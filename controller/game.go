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
			fmt.Println("START")
			game := models.InitGame(3)
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
			data := views.ActionRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)

			game := server.GetGame(data.Game)
			if err := game.PlayCardPublic(game.GetCurrentPlayer(), int(data.Details["position"].(float64)), data.Details["head"].(bool)); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(data)
			fmt.Println("END")
		}
	}
}
