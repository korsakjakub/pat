package pkg // import github.com/korsakjakub/pat/pkg

import "errors"

type Table struct {
	Players    []Player
	TableCards Deck
	Deck       Deck
	Settings   Settings
}

func NewTable(p []Player, s Settings) (*Table, error) {
	var err error
	if len(p) < 1 {
		err = errors.New("You need players in a table.")
	}
	return &Table{Players: p,
		TableCards: Deck{},
		Deck:       *NewDeck(),
		Settings:   s}, err
}
