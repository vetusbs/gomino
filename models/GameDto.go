package models

type GameDto struct {
	ID      string
	Cards   []CardDto
	Players []PlayerDto
}

func CreateGameDto(Game *Game) GameDto {
	current := Game.board.head
	cards := make([]CardDto, 0)
	for current != nil {
		cards = append(cards, CardDto{
			current.left,
			current.right,
		})
	}

	players := make([]PlayerDto, 0)

	for _, p := range Game.players {
		playerCards := make([]CardDto, len(p.cards))
		for i, p := range Game.players {
			playerCards[i] = CardDto{
				Left:  p.cards[i].left,
				Right: p.cards[i].right,
			}
		}
		players = append(players, PlayerDto{
			Cards:  playerCards,
			points: p.points,
		})
	}

	return GameDto{ID: Game.GetId(), Cards: cards, Players: players}
}

type PlayerDto struct {
	Cards  []CardDto
	points []int
}

type CardDto struct {
	Left  int
	Right int
}
