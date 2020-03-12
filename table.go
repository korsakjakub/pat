package pat

import "errors"

// Table - A structural representation of a physical Poker Table
type Table struct {
	Players    []Player
	TableCards Deck
	Deck       Deck
	Settings   Settings
}

// NewTable generates a new Table with at least one player in it
func NewTable(p []Player, s Settings) (*Table, error) {
	var err error
	if len(p) < 1 {
		err = errors.New("You need players in a table.")
	}
	return &Table{Players: p,
		TableCards: Deck{},
		Deck:       NewDeck(),
		Settings:   s}, err
}
