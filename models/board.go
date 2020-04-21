package models

import (
	"fmt"
)

// Board is bla
type Board struct {
	head *Card
	tail *Card
	sink []*Card
}

func (g *Board) PrintBoard() {
	for _, card := range g.sink {
		fmt.Printf("%v ", card.toString())
	}

	fmt.Printf("\nBOARD\n")
	var actual *Card
	actual = g.head
	for actual != nil {
		actual.Println()
		actual = actual.nextCard
	}
}
