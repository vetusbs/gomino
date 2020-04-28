package server

import (
	"github.com/vetusbs/gomino/models"
)

var hub Hub

type Hub struct {
	Games map[string]*models.Game
}

func InitHub() *Hub {
	hub = Hub{Games: make(map[string]*models.Game)}
	return &hub
}

func GetGame(id string) *models.Game {
	return hub.Games[id]
}

func AddGame(game *models.Game) {
	hub.Games[game.GetId()] = game
}
