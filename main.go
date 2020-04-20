package main

import (
	"fmt"
	"gomino/models"
)

func main() {
	game := models.InitGame(3)

	game.PrintGameState()

	for {
		var i int
		fmt.Scan(&i)
		fmt.Println("card ", i)
		game.Play(i)
		if i < 0 {
			break
		}
		game.PrintGameState()
	}
}
