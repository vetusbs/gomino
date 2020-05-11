package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type wsHub struct {
	connections  map[string]*websocket.Conn
	actionStream chan Game
}

func NewWsHub() wsHub {
	var wsHub = wsHub{}
	wsHub.connections = make(map[string]*websocket.Conn)
	wsHub.actionStream = make(chan Game)

	return wsHub
}

func (wsHub *wsHub) Run() {
	log.Println("Hub is running")
	for {
		select {
		case game := <-wsHub.actionStream:
			// Broadcast message exclude serverMessage.clientID
			gameString, _ := json.Marshal(game.Map())
			for _, connection := range wsHub.connections {
				connection.WriteMessage(websocket.TextMessage, []byte(gameString))
			}

			log.Println("Hub Broadcast message done " + game.id)
		}
		log.Println("HUB DONE")
	}
}

func (wsHub *wsHub) AddConnection(userId string, connection *websocket.Conn) {
	fmt.Printf("Add connection for user %v", userId)
	wsHub.connections[userId] = connection
}

func (wsHub *wsHub) SendMessage(userId string, message string) {
	if wsHub.connections[userId] != nil {
		wsHub.connections[userId].WriteMessage(websocket.TextMessage, []byte(message))
	}
}

func (wsHub *wsHub) broadCast(game Game) {
	wsHub.actionStream <- game
}
