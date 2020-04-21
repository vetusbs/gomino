package models

import (
	"fmt"
	"math/rand"
	"time"
)

// Game is bla
type Game struct {
	board         *Board
	players       []*Player
	currentPlayer int
	winner        *Player
}

func (game *Game) GetCurrentPlayer() *Player {
	return game.players[game.currentPlayer]
}

func (game *Game) validatePlayHead(card *Card) bool {
	if game.board.head.reverse == true {
		if game.board.head.right == card.left || game.board.head.right == card.right {
			return true
		}
	} else {
		if game.board.head.left == card.left || game.board.head.left == card.right {
			return true
		}
	}
	return false
}

func (game *Game) validatePlayTail(card *Card) bool {
	if game.board.tail.reverse == true {
		if game.board.tail.left == card.left || game.board.tail.left == card.right {
			return true
		}
	} else {
		if game.board.tail.right == card.left || game.board.tail.right == card.right {
			return true
		}
	}
	return false
}

func (game *Game) playCard(card *Card) bool {
	if game.board.head == nil {
		game.board.head = card
		game.board.tail = card
	} else if game.validatePlayHead(card) {
		if game.board.tail.getFreeNumber() == card.left {
			card.reverse = true
		}
		game.board.head.prevCard = card
		card.nextCard = game.board.head
		game.board.head = card
	} else if game.validatePlayTail(card) {
		if game.board.tail.getFreeNumber() == card.right {
			card.reverse = true
		}
		game.board.tail.nextCard = card
		card.prevCard = game.board.tail
		game.board.tail = card
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

	return Game{
		players:       players,
		currentPlayer: 0,
		board: &Board{
			sink: gameCards[numberOfPlayers*nCardsPerUser : 28],
		},
	}
}

func (game *Game) PrintGameState() {
	fmt.Printf("\nGAME STATUS\nSINK\n")

	game.board.PrintBoard()

	fmt.Printf("\nPLAYERS\n")
	for _, p := range game.players {
		p.Println()
	}
}
