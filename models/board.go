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

func (board *Board) playCard(card *Card) bool {
	// TODO: fix the reverse for the first card, when left and right are nil.
	if board.head == nil {
		board.head = card
		board.tail = card
	} else if board.validatePlayHead(card) {
		fmt.Printf("validated for head %v --- %v \n", card, board.head)
		if board.head.getFreeNumber() == card.left {
			fmt.Println("set reverse to true")
			card.reverse = true
		}
		board.head.prevCard = card
		card.nextCard = board.head
		board.head = card
	} else if board.validatePlayTail(card) {
		fmt.Println("validated for tail")
		if board.tail.getFreeNumber() == card.right {
			fmt.Println("set reverse to true")
			card.reverse = true
		}
		board.tail.nextCard = card
		card.prevCard = board.tail
		board.tail = card
	} else {
		return false
	}

	return true
}

func (board *Board) validatePlayHead(card *Card) bool {
	if board.head.reverse == true {
		fmt.Println("valid for reverse")
		if board.head.right == card.left || board.head.right == card.right {
			return true
		}
	} else {
		fmt.Println("valid for not reverse")
		if board.head.left == card.left || board.head.left == card.right {
			return true
		}
	}
	return false
}

func (board *Board) validatePlayTail(card *Card) bool {
	if board.tail.reverse == true {
		fmt.Println("valid for reverse")
		if board.tail.left == card.left || board.tail.left == card.right {
			return true
		}
	} else {
		fmt.Println("valid for not reverse")
		if board.tail.right == card.left || board.tail.right == card.right {
			return true
		}
	}
	return false
}
