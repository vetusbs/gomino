package models

import (
	"fmt"
	"math/rand"
	"time"
)

// Game is bla
type Game struct {
	head          *Card
	tail          *Card
	sink          []*Card
	players       []*Player
	currentPlayer int
}

func (game *Game) Play(position int) {
	card := game.players[game.currentPlayer].cards[position]
	if game.head == nil {
		game.head = card
		game.tail = card
	} else {
		game.head.leftCard = card
		card.rightCard = game.head
		game.head = card
	}
	game.players[game.currentPlayer].play(position)
}

func (game *Game) playCard(card *Card) {
	if game.head == nil {
		game.head = card
		game.tail = card
	} else {
		game.head.leftCard = card
		card.rightCard = game.head
		game.head = card
	}
}

// InitGame Creates an empty game.
func InitGame(numberOfPlayers int) Game {
	gameCards := make([]*Card, 28)
	total := 0
	for i := 0; i < 7; i++ {
		for j := i; j < 7; j++ {
			gameCards[total] = &Card{left: i, right: j}
			total++
		}
	}

	players := make([]*Player, numberOfPlayers)

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
		players[i] = &Player{
			name:  fmt.Sprint("player-", i),
			cards: playerCards,
		}
	}

	return Game{sink: gameCards[numberOfPlayers*nCardsPerUser : 28], players: players, currentPlayer: 0}
}

func (g *Game) PrintGameState() {
	fmt.Printf("\nGAME STATUS\nSINK\n")

	for _, p := range g.sink {
		fmt.Printf("%v ", p.toString())
	}

	fmt.Printf("\nBOARD\n")
	var actual *Card
	actual = g.head
	for actual != nil {
		actual.Println()
		actual = actual.rightCard
	}
	fmt.Printf("\nPLAYERS\n")
	for _, p := range g.players {
		p.Println()
	}
}
