package models

import (
	"github.com/vetusbs/gomino/controller/dto"
)

func (Game Game) Map() dto.GameDto {
	current := Game.board.head
	cards := make([]dto.CardDto, 0)
	for current != nil {
		cards = append(cards, dto.CardDto{
			Left:    current.left,
			Right:   current.right,
			Reverse: current.reverse,
		})
		current = current.nextCard
	}

	sink := make([]dto.CardDto, 0)
	for _, cardSink := range Game.board.sink {
		sink = append(sink, dto.CardDto{
			Left:    cardSink.left,
			Right:   cardSink.right,
			Reverse: cardSink.reverse,
		})
	}

	players := make([]dto.PlayerDto, 0)

	for _, player := range Game.players {
		playerCards := make([]dto.CardDto, len(player.cards))
		for i, card := range player.cards {
			playerCards[i] = dto.CardDto{
				Left:    card.left,
				Right:   card.right,
				Reverse: card.reverse,
			}
		}

		players = append(players, dto.PlayerDto{
			Cards:           playerCards,
			Points:          player.points,
			Name:            player.GetName(),
			IsCurrentPlayer: player == Game.GetCurrentPlayer(),
		})
	}

	Game.PrintGameState()

	return dto.GameDto{ID: Game.GetId(), Cards: cards, Players: players, Sink: len(Game.board.sink)}
}
