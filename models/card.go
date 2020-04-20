package models

import (
	"fmt"
)

func init() {
	fmt.Println("init in sandbox.go")
}

// Card is bla
type Card struct {
	left      int
	right     int
	leftCard  *Card
	rightCard *Card
	reverse   bool
}

func (c Card) Println() {
	if c.reverse {
		fmt.Println(c.left, "·", c.right)
	} else {
		fmt.Println(c.right, "·", c.left)
	}

}

func (c Card) isDouble() bool {
	return c.left == c.right
}
