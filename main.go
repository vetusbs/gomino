package main

import (
	"domino/domino/controller"
	"domino/domino/models"
	"fmt"
)

func main() {
	fmt.Println("Hello world")

	game := models.InitGame(3)

	fmt.Println(game)

	game.PrintGameState()
	controller.Ping()
}
