package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type wsHub struct {
	clients      map[string]wsClient
	actionStream chan Game
}

func NewWsHub() wsHub {
	var wsHub = wsHub{}
	wsHub.clients = make(map[string]wsClient)
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
			for _, wsClient := range wsHub.clients {
				wsClient.send <- gameString
			}

			log.Println("Hub Broadcast message done " + game.id)
		}
		log.Println("HUB DONE")
	}
}

func (wsHub *wsHub) AddConnection(userId string, connection *websocket.Conn) {
	fmt.Printf("Add connection for user %v", userId)
	wsClient := NewClient(connection, *wsHub)

	go wsClient.ReadPump()
	go wsClient.WritePump()

	wsHub.clients[userId] = wsClient
}

func (wsHub *wsHub) broadCast(game Game) {
	logrus.Debugf(" broadcasting ")
	wsHub.actionStream <- game
}
