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
		fmt.Printf("%v: ", currentPlayer.GetName())
		fmt.Scan(&i)
		fmt.Printf("%v %v\n", currentPlayer.GetName(), i)

		if i == 33 {
			currentPlayer.AutoPlay(&game)
		} else {
			if game.PlayCardPublic(currentPlayer, i) == false {
				game.Pick(currentPlayer)
			}
		}

		if i < 0 {
			break
		}

		game.PrintGameState()
		fmt.Println("*******************************")
	}
}
