package main

import (
	"fmt"

	"github.com/vetusbs/gomino/models"
)

func main() {
	game := models.InitGame(3)

	game.PrintGameState()
	var currentPlayer *models.Player
	for game.IsFinished() == false {
		currentPlayer = game.GetCurrentPlayer()
		var i int
		var head bool
		fmt.Printf("%v: ", currentPlayer.GetName())
		fmt.Scan(&i)

		if i == 33 {
			currentPlayer.AutoPlay(&game)
		} else {
			fmt.Scan(&head)
			fmt.Printf("%v %v %v\n", currentPlayer.GetName(), i, head)
			if err := game.PlayCardPublic(currentPlayer, i, head); err != nil {
				fmt.Println(err)
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
