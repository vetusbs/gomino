package dto

type GameDto struct {
	ID      string      `json:"id"`
	Cards   []CardDto   `json:"cards"`
	Players []PlayerDto `json:"players"`
	Sink    int         `json:"sink"`
}

type PlayerDto struct {
	Cards           []CardDto `json:"cards"`
	Points          []int     `json:"points"`
	Name            string    `json:"name"`
	IsCurrentPlayer bool      `json:"isCurrentPlayer"`
}

type CardDto struct {
	Left    int  `json:"left"`
	Right   int  `json:"right"`
	Reverse bool `json:"reverse"`
}
