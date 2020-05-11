package models

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Card is a representation of one of the 28 cards of the domino game.
// we assume that the first number is the left, and the second is the right
// if the card is reversed, it means that the card shoult be displayed from right to left
type Card struct {
	left     int
	right    int
	prevCard *Card
	nextCard *Card
	reverse  bool
}

func (c *Card) toString() string {
	if c.reverse == false {
		return fmt.Sprintf("[%d|%d]", c.left, c.right)
	} else {
		return fmt.Sprintf("[%d|%d]", c.right, c.left)
	}
}

func (c Card) Println() {
	if c.reverse == false {
		log.Debugf("[%d|%d] ", c.left, c.right)
	} else {
		log.Debugf("[%d|%d] ", c.right, c.left)
	}
}

func (c Card) getPoints() int {
	return c.left + c.right
}

func (c Card) isDouble() bool {
	return c.left == c.right
}

func (c Card) isInitialCard() bool {
	return c.prevCard == nil && c.nextCard == nil
}

func (c Card) getFreeNumber() int {
	if c.reverse == false {
		if c.nextCard == nil {
			return c.right
		} else {
			return c.left
		}
	} else {
		if c.nextCard == nil {
			return c.left
		} else {
			return c.right
		}
	}
}
