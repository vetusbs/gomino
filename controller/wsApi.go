package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vetusbs/gomino/server"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type msg struct {
	Num int
}

func main2() {
	http.HandleFunc("/", rootHandler)

	panic(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Root path")
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == http.MethodOptions {
		return
	}

	fmt.Println("START")
	//userID := mux.Vars(r)["userId"]
	if r.Header.Get("Origin") != "http://"+r.Host {
		//http.Error(w, "Origin not allowed", 403)
		//		return
	}
	fmt.Println("Request received")
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	userID := mux.Vars(r)["userId"]
	gameID := mux.Vars(r)["gameId"]

	game := server.GetGame(gameID)
	game.AddConnection(userID, conn)

	//go echo(conn)
}

func echo(conn *websocket.Conn) {
	for {
		m := msg{}

		err := conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}

		fmt.Printf("Got message: %#v\n", m)

		//if err = conn.WriteJSON(m); err != nil {
		//	fmt.Println(err)
		//}
	}
}
