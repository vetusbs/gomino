package models

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("init in sandbox.go")
}

// Game is bla
type Game struct {
	head    *Card
	tail    *Card
	sink    []Card
	Players []Player
}

// InitGame Creates an empty game.
func InitGame(numberOfPlayers int) Game {
	gameCards := make([]Card, 28)
	total := 0
	for i := 0; i < 7; i++ {
		for j := i; j < 7; j++ {
			gameCards[total] = Card{left: i, right: j}
			total++
		}
	}

	players := make([]Player, numberOfPlayers)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(gameCards), func(i, j int) { gameCards[i], gameCards[j] = gameCards[j], gameCards[i] })

	var nCardsPerUser int
	if numberOfPlayers == 2 {
		nCardsPerUser = 14
	} else if numberOfPlayers == 3 {
		nCardsPerUser = 7
	} else if numberOfPlayers == 4 {
		nCardsPerUser = 7
	}

	for i := 0; i < numberOfPlayers; i++ {
		playerCards := gameCards[i*nCardsPerUser : i*nCardsPerUser+nCardsPerUser]
		players[i] = Player{
			name:  fmt.Sprint("player-", i),
			cards: playerCards,
		}
	}

	return Game{sink: gameCards[numberOfPlayers*nCardsPerUser : 28], Players: players}
}

func (g Game) PrintGameState() {
	fmt.Println("SINK")
	for _, p := range g.sink {
		p.Println()
	}

	fmt.Println("GAME STATUS")
	var actual *Card
	actual = g.head
	for actual != nil {
		actual.Println()
		actual = actual.rightCard
	}

	for _, p := range g.Players {
		p.Println()
	}
}
