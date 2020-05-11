package models

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 512
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

func NewClient(conn *websocket.Conn, wsHub wsHub) wsClient {
	wsClient := wsClient{}

	wsClient.wsHub = wsHub
	wsClient.conn = conn
	wsClient.send = make(chan []byte)

	return wsClient
}

func (wsClient *wsClient) ReadPump() {
	defer func() {
		log.Println("Close readpump", wsClient.id)
		//wsClient.wsHub.UnRegister(wsClient)
		wsClient.conn.Close()
	}()
	for {
		log.Println("Waiting for message")
		_, message, err := wsClient.conn.ReadMessage()
		if err != nil {
			// Client disconnect
			log.Println(message)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
	}
}

func (wsClient *wsClient) WritePump() {
	defer func() {
		wsClient.conn.Close()
	}()

	for {
		select {
		case message, ok := <-wsClient.send:
			fmt.Println("send message")
			// NOTE: if there is remaining in send, will cause deadlock
			wsClient.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				wsClient.conn.WriteMessage(websocket.CloseMessage, []byte{})
				fmt.Println("Write pump closed", wsClient.id)
				return
			}

			err := wsClient.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				fmt.Println("Write pump cannot closed", wsClient.id, err)
				return
			}
			/*
					w, err := wsClient.conn.NextWriter(websocket.BinaryMessage)
					if err != nil {
						fmt.Println("Write pump closed ", wsClient.id, err)
						return
					}
					w.Write(message)

				// Add queued chat messages to the current websocket message.
				if err := w.Close(); err != nil {
					fmt.Println("Write pump cannot closed", wsClient.id, err)
					return
				}
			*/
		}
	}
}
