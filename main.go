package main

import (
	"fmt"
	"gomino/models"
)

func main() {
	game := models.InitGame(3)

	game.PrintGameState()
	var currentPlayer *models.Player
	for {
		currentPlayer = game.GetCurrentPlayer()
		var i int
		fmt.Scan(&i)
		fmt.Printf("%v %v\n", currentPlayer.GetName(), i)
		if currentPlayer.Play(i, &game) == false {
			currentPlayer.Pick(0, &game)
		}
		if i < 0 {
			break
		}

		game.PrintGameState()
		fmt.Println("*******************************")
	}
}
