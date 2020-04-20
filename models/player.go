package models

import (
	"fmt"
)

func init() {
	fmt.Println("init in sandbox.go")
}

// Player is bla
type Player struct {
	name  string
	cards []Card
}

func (p Player) Println() {
	fmt.Println(p.name)
	for i := 0; i < len(p.cards); i++ {
		p.cards[i].Println()
	}
}
