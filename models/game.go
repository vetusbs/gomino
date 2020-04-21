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
	winner        *Player
}

func (game *Game) GetCurrentPlayer() *Player {
	return game.players[game.currentPlayer]
}

func (game *Game) validatePlayHead(card *Card) bool {
	if game.head.reverse == true {
		fmt.Println("valid for reverse")
		if game.head.right == card.left || game.head.right == card.right {
			return true
		}
	} else {
		fmt.Println("valid for not reverse")
		if game.head.left == card.left || game.head.left == card.right {
			return true
		}
	}
	return false
}

func (game *Game) validatePlayTail(card *Card) bool {
	if game.tail.reverse == true {
		fmt.Println("valid for reverse")
		if game.tail.left == card.left || game.tail.left == card.right {
			return true
		}
	} else {
		fmt.Println("valid for not reverse")
		if game.tail.right == card.left || game.tail.right == card.right {
			return true
		}
	}
	return false
}

func (game *Game) playCard(card *Card) bool {
	if game.head == nil {
		game.head = card
		game.tail = card
	} else if game.validatePlayHead(card) {
		if game.tail.getFreeNumber() == card.left {
			card.reverse = true
		}
		game.head.prevCard = card
		card.nextCard = game.head
		game.head = card
	} else if game.validatePlayTail(card) {
		if game.tail.getFreeNumber() == card.right {
			card.reverse = true
		}
		game.tail.nextCard = card
		card.prevCard = game.tail
		game.tail = card
	} else {
		game.currentPlayer = (game.currentPlayer + 1) % len(game.players)
		return false
	}
	if len(game.players[game.currentPlayer].cards) == 0 {
		fmt.Printf("\n*****************************************")
		fmt.Printf("\n***** GAME IS OVER %v WINS ***", game.players[game.currentPlayer].name)
		fmt.Printf("\n*****************************************\n")
	}
	game.currentPlayer = (game.currentPlayer + 1) % len(game.players)

	return true
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
	fmt.Printf("\nPLAYERS\n")
	for _, p := range g.players {
		p.Println()
	}
}
