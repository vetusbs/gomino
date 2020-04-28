package models

import "fmt"

type GameDto struct {
	ID      string `json:"id"`
	Cards   []CardDto
	Players []PlayerDto
}

func CreateGameDto(Game *Game) GameDto {
	fmt.Println("CREATE GAME_DTO")
	current := Game.board.head
	cards := make([]CardDto, 0)
	for current != nil {
		cards = append(cards, CardDto{
			current.left,
			current.right,
		})
		current = current.nextCard
	}
	fmt.Println("CREATE GAME_DTO END BOARD")
	players := make([]PlayerDto, 0)

	for _, player := range Game.players {
		playerCards := make([]CardDto, len(player.cards))
		for i, card := range player.cards {
			playerCards[i] = CardDto{
				Left:  card.left,
				Right: card.right,
			}
		}
		players = append(players, PlayerDto{
			Cards:  playerCards,
			points: player.points,
			Name:   player.GetName(),
		})
	}

	fmt.Println("CREATE GAME_DTO END PLAYERS")
	Game.PrintGameState()

	return GameDto{ID: Game.GetId(), Cards: cards, Players: players}
}

type PlayerDto struct {
	Cards  []CardDto
	points []int
	Name   string
}

type CardDto struct {
	Left  int `json:"left"`
	Right int `json:"right"`
}
