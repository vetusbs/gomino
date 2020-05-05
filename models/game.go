package models

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/vetusbs/gomino/views"

	"github.com/google/uuid"
)

// Game
type Game struct {
	id            string
	board         *Board
	players       []*Player
	currentPlayer int
}

func (game *Game) GetId() string { return game.id }

func (game *Game) IsFinished() bool {
	for _, p := range game.players {
		if p.getSumOfPoints() > 100 {
			return true
		}
	}
	return false
}

func (game *Game) GetCurrentPlayer() *Player {
	return game.players[game.currentPlayer]
}

func (game *Game) nextPlayer() {
	game.currentPlayer = (game.currentPlayer + 1) % len(game.players)
}

func (game *Game) PlayCard(player *Player, inputCard CardDto, isLeft bool) error {
	for i, card := range player.cards {
		fmt.Println(card, inputCard)
		if card.left == inputCard.Left && card.right == inputCard.Right {
			return game.PlayCardPublic(player, i, isLeft)
		}
	}
	return fmt.Errorf("player: %v does not have this card", player.name)
}

func (game *Game) PlayCardPublic(player *Player, cardPosition int, isLeft bool) error {

	if player != game.players[game.currentPlayer] {
		return errors.New("this is not the current player")
	}

	if cardPosition >= len(player.cards) {
		return errors.New("this is not a valid position")
	}

	card := player.cards[cardPosition]
	result := game.board.playCard(card, isLeft)
	if result == nil {
		player.play(cardPosition)
		if len(game.players[game.currentPlayer].cards) == 0 {
			fmt.Printf("\n*****************************************")
			fmt.Printf("\n***** GAME IS OVER %v WINS ***", game.players[game.currentPlayer].name)
			fmt.Printf("\n*****************************************\n")
			game.addPoints()
			game.restartGame()
		} else {
			game.nextPlayer()
		}
		return nil
	}
	return result
}

func (game *Game) AddUser(userId string, userName string) error {
	for _, player := range game.players {
		if player.userID == "" {
			player.name = userName
			player.userID = userId
			return nil
		}
	}

	return errors.New("There are no cards to pick")
}

func (game *Game) Pick(player *Player) error {
	if game.players[game.currentPlayer] == player {
		if len(game.board.sink) > 0 {
			player.pick(0, game)
			return nil
		}
	}
	game.nextPlayer()
	return errors.New("There are no cards to pick")
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

func (game *Game) restartGame() {
	gameCards := createCards()
	nCardsPerUser := cardsPerUser(len(game.players))

	for i := 0; i < len(game.players); i++ {
		playerCards := gameCards[i*nCardsPerUser : i*nCardsPerUser+nCardsPerUser]
		game.players[i].cards = playerCards
	}
	game.board = &Board{
		sink: gameCards[len(game.players)*nCardsPerUser : 28],
	}
}

// InitGame Creates an empty game.
func InitGame(createGameRequest views.CreateGameRequest) Game {

	numberOfPlayers := func() int {
		if createGameRequest.Players == 0 {
			return 4
		} else {
			return createGameRequest.Players
		}
	}()

	gameId := func() string {
		if createGameRequest.ID == "" {
			uuid, _ := uuid.NewUUID()
			return uuid.String()
		} else {
			return createGameRequest.ID
		}
	}()

	gameCards := createCards()
	players := make([]*Player, numberOfPlayers)
	nCardsPerUser := cardsPerUser(numberOfPlayers)

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
		id:            gameId,
		players:       players,
		currentPlayer: 0,
		board: &Board{
			sink: gameCards[numberOfPlayers*nCardsPerUser : 28],
		},
	}
}

func createCards() []*Card {
	gameCards := make([]*Card, 28)
	total := 0
	for i := 0; i < 7; i++ {
		for j := i; j < 7; j++ {
			gameCards[total] = &Card{left: i, right: j}
			total++
		}
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(gameCards), func(i, j int) { gameCards[i], gameCards[j] = gameCards[j], gameCards[i] })

	return gameCards
}

func cardsPerUser(numberOfPlayers int) int {
	if numberOfPlayers == 2 {
		return 7
	} else if numberOfPlayers == 3 {
		return 7
	} else if numberOfPlayers == 4 {
		return 7
	}
	return 5
}

func (game *Game) PrintGameState() {
	fmt.Printf("\nGAME STATUS\nSINK\n")

	game.board.PrintBoard()

	fmt.Printf("\nPLAYERS\n")
	for _, p := range game.players {
		p.Println()
	}
}
