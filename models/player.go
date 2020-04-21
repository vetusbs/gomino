package models

import (
	"fmt"
)

// Player is bla
type Player struct {
	name  string
	cards []*Card
}

func (player *Player) GetName() string {
	return player.name
}

// Println prints player
func (player Player) Println() {
	fmt.Print(player.name, " -- ", len(player.cards), " -- ")
	for _, card := range player.cards {
		//fmt.Printf("%+v\n", card)
		card.Println()
	}
	fmt.Println()
}

func (player *Player) Play(position int, game *Game) bool {
	card := player.cards[position]
	if game.playCard(card) == true {
		player.cards = remove(player.cards, position)
		return true
	} else {
		fmt.Printf("\nyou cannot play this card\n")
		return false
	}
}

func (player *Player) Pick(position int, game *Game) {
	card := game.board.sink[position]
	game.board.sink = remove(game.board.sink, position)
	//	player.cards = append(player.cards, card)
	lalala := append(make([]*Card, 0), player.cards...)
	lalala = append(lalala, card)
	player.cards = lalala
}

func remove(slice []*Card, i int) []*Card {
	// Remove the element at index i from a.
	copy(slice[i:], slice[i+1:]) // Shift a[i+1:] left one index.
	slice[len(slice)-1] = nil    // Erase last element (write zero value).
	slice = slice[:len(slice)-1] // Truncate slice.

	return slice
}
