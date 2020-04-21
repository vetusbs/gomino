package models

import (
	"fmt"
)

// Player is bla
type Player struct {
	name  string
	cards []*Card
}

// Println prints player
func (player Player) Println() {
	fmt.Print(player.name, " -- ", len(player.cards), " -- ")
	for i := 0; i < len(player.cards); i++ {
		player.cards[i].Println()
	}
	fmt.Println()
}

func (player *Player) Play(position int, game *Game) bool {
	card := player.cards[position]
	if game.playCard(card) {
		player.cards = remove(player.cards, position)
		return true
	} else {
		fmt.Printf("\nyou cannot play this card\n")
		return false
	}
}

func (player *Player) Pick(position int, game *Game) {
	card := game.sink[position]
	game.sink = remove(game.sink, position)
	player.cards = append(player.cards, card)
}

func remove(slice []*Card, i int) []*Card {
	// Remove the element at index i from a.
	copy(slice[i:], slice[i+1:]) // Shift a[i+1:] left one index.
	slice[len(slice)-1] = nil    // Erase last element (write zero value).
	slice = slice[:len(slice)-1] // Truncate slice.

	return slice
}
