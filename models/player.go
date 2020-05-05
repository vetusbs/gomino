package models

import (
	"errors"
	"fmt"
)

// Player is bla
type Player struct {
	name   string
	cards  []*Card
	points []int
	userID string
}

func NewPlayer(name string, cards []*Card, points []int) Player {
	return Player{
		name:   name,
		cards:  cards,
		points: points,
	}
}

func (player *Player) GetName() string {
	return player.name
}

// Println prints player
func (player Player) Println() {
	fmt.Print(player.name, " -- ", len(player.cards), " -- ")
	for _, card := range player.cards {
		card.Println()
	}
	fmt.Printf(" points: %v\n", player.points)
}

func (player *Player) play(position int) error {
	if len(player.cards) == 0 {
		return errors.New("This player has no cards mama")
	}
	if position > len(player.cards) {
		return errors.New("this position does not exist")
	}

	player.cards = remove(player.cards, position)

	return nil
}

func (player *Player) AutoPlay(game *Game) error {
	for position, _ := range player.cards {
		if game.PlayCardPublic(player, position, true) == nil {
			return nil
		}
		if game.PlayCardPublic(player, position, false) == nil {
			return nil
		}
	}
	if err := game.Pick(player); err == nil {
		return player.AutoPlay(game)
	}
	return errors.New("There is no possible card to play")
}

func (player *Player) getSumOfPoints() int {
	sum := 0
	for _, value := range player.points {
		sum = sum + value
	}
	return sum
}

func (player *Player) pick(position int, game *Game) {
	card := game.board.sink[position]
	game.board.sink = remove(game.board.sink, position)
	//	player.cards = append(player.cards, card)
	lalala := append(make([]*Card, 0), player.cards...)
	lalala = append(lalala, card)
	player.cards = lalala
}

func (player *Player) countPoints() int {
	sum := 0
	for _, c := range player.cards {
		sum = sum + c.getPoints()
	}
	return sum
}

func remove(slice []*Card, i int) []*Card {
	// Remove the element at index i from a.
	copy(slice[i:], slice[i+1:]) // Shift a[i+1:] left one index.
	slice[len(slice)-1] = nil    // Erase last element (write zero value).
	slice = slice[:len(slice)-1] // Truncate slice.

	return slice
}
