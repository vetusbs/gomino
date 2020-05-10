package server

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var wsHub WsHub

func init() {
	fmt.Println("init wshub")
	wsHub = WsHub{connections: make(map[string]*websocket.Conn)}
}

type WsHub struct {
	connections map[string]*websocket.Conn
}

func AddConnection(userId string, connection *websocket.Conn) {
	fmt.Printf("Add connection for user %v", userId)
	wsHub.connections[userId] = connection
}

func SendMessage(userId string, message string) {
	if wsHub.connections[userId] != nil {
		wsHub.connections[userId].WriteMessage(websocket.TextMessage, []byte(message))
	}
}
