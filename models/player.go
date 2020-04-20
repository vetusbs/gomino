package models

import (
	"fmt"
)

// Player is bla
type Player struct {
	name  string
	cards []*Card
}

func (p Player) Println() {
	fmt.Print(p.name, " -- ", len(p.cards), " -- ")
	for i := 0; i < len(p.cards); i++ {
		p.cards[i].Println()
	}
	fmt.Println()
}

func (player *Player) play(position int) {
	player.cards = remove(player.cards, position)
}

func remove(slice []*Card, i int) []*Card {
	// Remove the element at index i from a.
	copy(slice[i:], slice[i+1:]) // Shift a[i+1:] left one index.
	slice[len(slice)-1] = nil    // Erase last element (write zero value).
	slice = slice[:len(slice)-1] // Truncate slice.

	return slice
}
