package models

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

// Board is bla
type Board struct {
	head *Card
	tail *Card
	sink []*Card
}

func (g *Board) PrintBoard() {
	for _, card := range g.sink {
		log.Debug("%v ", card.toString())
	}

	log.Debugf("\nBOARD\n")
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
// isLeft : true if the card number is the left
func (board *Board) playCard(card *Card, isLeft bool) error {
	// TODO: fix the reverse for the first card, when left and right are nil.
	var targetNumber int
	if isLeft == true {
		targetNumber = card.left
	} else {
		targetNumber = card.right
	}
	if board.head == nil {
		board.head = card
		board.tail = card
	} else if board.validatePlayHead(targetNumber) {
		if board.head.isInitialCard() && card.right != board.head.left {
			card.reverse = true
		} else if board.head.getFreeNumber() == card.left {
			card.reverse = true
		}
		board.head.prevCard = card
		card.nextCard = board.head
		board.head = card
	} else if board.validatePlayTail(targetNumber) {
		if board.head.isInitialCard() && card.left != board.head.right {
			card.reverse = true
		} else if board.tail.getFreeNumber() == card.right {
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

func (board *Board) validatePlayHead(targetNumber int) bool {
	if board.head.reverse == true {
		if board.head.right == targetNumber {
			return true
		}
	} else {
		if board.head.left == targetNumber {
			return true
		}
	}
	return false
}

func (board *Board) validatePlayTail(targetNumber int) bool {
	if board.tail.reverse == true {
		if board.tail.left == targetNumber {
			return true
		}
	} else {
		if board.tail.right == targetNumber {
			return true
		}
	}
	return false
}

func (board *Board) isItClosed() (bool, error) {
	counter := 0
	if board.head != nil && board.head.getFreeNumber() == board.tail.getFreeNumber() {
		current := board.head

		for current != nil {
			log.Debugf(" current values: %v ", counter)
			if current.left == board.head.getFreeNumber() {
				counter++
			}

			if current.right == board.head.getFreeNumber() {
				counter++
			}

			current = current.nextCard
		}
	}
	return counter == 8, nil
}
