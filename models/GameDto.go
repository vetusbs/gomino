package models

import "fmt"

type GameDto struct {
	ID      string      `json:"id"`
	Cards   []CardDto   `json:"cards"`
	Players []PlayerDto `json:"players"`
	Sink    int         `json:"sink"`
}

func CreateGameDto(Game *Game) GameDto {
	fmt.Println("CREATE GAME_DTO")
	current := Game.board.head
	cards := make([]CardDto, 0)
	for current != nil {
		cards = append(cards, CardDto{
			Left:    current.left,
			Right:   current.right,
			Reverse: current.reverse,
		})
		current = current.nextCard
	}

	sink := make([]CardDto, 0)
	for _, cardSink := range Game.board.sink {
		sink = append(sink, CardDto{
			Left:    cardSink.left,
			Right:   cardSink.right,
			Reverse: cardSink.reverse,
		})
	}

	fmt.Println("CREATE GAME_DTO END BOARD")
	players := make([]PlayerDto, 0)

	for _, player := range Game.players {
		playerCards := make([]CardDto, len(player.cards))
		for i, card := range player.cards {
			playerCards[i] = CardDto{
				Left:    card.left,
				Right:   card.right,
				Reverse: card.reverse,
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

	return GameDto{ID: Game.GetId(), Cards: cards, Players: players, Sink: len(Game.board.sink)}
}

type PlayerDto struct {
	Cards  []CardDto
	points []int
	Name   string
}

type CardDto struct {
	Left    int  `json:"left"`
	Right   int  `json:"right"`
	Reverse bool `json:"reverse"`
}
