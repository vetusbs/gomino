package models

import (
	"net"

	"github.com/gorilla/websocket"
)

type wsClient struct {
	wsHub wsHub
	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	// ID
	id int32

	// RemoteAddress
	remoteAdd net.Addr
}
