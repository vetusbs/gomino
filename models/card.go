package models

import (
	"fmt"
)

// Card is bla
type Card struct {
	left      int
	right     int
	leftCard  *Card
	rightCard *Card
	reverse   bool
}

func (c *Card) toString() string {
	return fmt.Sprintf("%d\u00B7%d", c.left, c.right)
}

func (c Card) Println() {
	if c.reverse {
		fmt.Printf("%d\u00B7%d ", c.left, c.right)
	} else {
		fmt.Printf("%d\u00B7%d ", c.right, c.left)
	}
}

func (c Card) isDouble() bool {
	return c.left == c.right
}
