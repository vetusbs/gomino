package models

import (
	"errors"
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

// playCard is adding the actual card into the board, and make sure that
// the movemnt is allowed.
// params:
// card : the actual card to put into the board
// head : true if the card goes to the head. false if it goes to the tail.
func (board *Board) playCard(card *Card, head bool) error {
	// TODO: fix the reverse for the first card, when left and right are nil.
	if board.head == nil {
		board.head = card
		board.tail = card
	} else if head == true && board.validatePlayHead(card) {
		fmt.Printf("validated for head %v --- %v \n", card, board.head)
		if board.head.isInitialCard() && card.right != board.head.left {
			card.reverse = true
		} else if board.head.getFreeNumber() == card.left {
			fmt.Println("set reverse to true")
			card.reverse = true
		}
		board.head.prevCard = card
		card.nextCard = board.head
		board.head = card
	} else if head == false && board.validatePlayTail(card) {
		fmt.Println("validated for tail")
		if board.head.isInitialCard() && card.left != board.head.right {
			card.reverse = true
		} else if board.tail.getFreeNumber() == card.right {
			fmt.Println("set reverse to true")
			card.reverse = true
		}
		board.tail.nextCard = card
		card.prevCard = board.tail
		board.tail = card
	} else {
		return errors.New("You cannot put this card")
	}

	return nil
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
