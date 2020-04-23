package models

import (
	"fmt"
	"math/rand"
	"time"
)

// Game
type Game struct {
	board         *Board
	players       []*Player
	currentPlayer int
}

func (game *Game) IsFinished() bool {
	for _, p := range game.players {
		if p.getSumOfPoints() > 10 {
			return true
		}
	}
	return false
}

func (game *Game) GetCurrentPlayer() *Player {
	return game.players[game.currentPlayer]
}

func (game *Game) playCard(card *Card) bool {

	result := game.board.playCard(card)
	if result == true {
		if len(game.players[game.currentPlayer].cards) == 0 {
			fmt.Printf("\n*****************************************")
			fmt.Printf("\n***** GAME IS OVER %v WINS ***", game.players[game.currentPlayer].name)
			fmt.Printf("\n*****************************************\n")
			game.addPoints()
		}
		game.nextPlayer()
	}

	return result
}

func (game *Game) nextPlayer() {
	game.currentPlayer = (game.currentPlayer + 1) % len(game.players)
}

func (game *Game) PlayCardPublic(player *Player, cardPosition int) bool {

	if player != game.players[game.currentPlayer] {
		return false
	}

	card := player.cards[cardPosition]
	result := game.board.playCard(card)
	if result == true {
		player.play(cardPosition)
		if len(game.players[game.currentPlayer].cards) == 0 {
			fmt.Printf("\n*****************************************")
			fmt.Printf("\n***** GAME IS OVER %v WINS ***", game.players[game.currentPlayer].name)
			fmt.Printf("\n*****************************************\n")
			game.addPoints()
		}
		game.currentPlayer = (game.currentPlayer + 1) % len(game.players)
	}

	return result
}

func (game *Game) Pick(player *Player) bool {
	if game.players[game.currentPlayer] == player {
		if len(game.board.sink) > 0 {
			player.pick(0, game)
			return true
		}
	}
	game.nextPlayer()
	return false
}

func (game *Game) addPoints() {
	for _, player := range game.players {
		if player == game.players[game.currentPlayer] {
			player.points = append(player.points, 0)
		} else {
			player.points = append(player.points, player.countPoints())
		}
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

	//cardtest := game.board.sink[0]
	//game.board.sink = remove(game.board.sink, 0)
	//game.players[1].cards = append(game.players[1].cards, cardtest)

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
