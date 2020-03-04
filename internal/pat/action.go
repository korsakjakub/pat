package pat

import (
	"errors"
	"fmt"
)

// Action: Possible actions for a Player
type Action int

// Fold - folds (give out cards), Check - wait for the next card, but don't bring in new chips, Raise - wait for the next card, and bring in new chips.
const (
	Fold Action = iota
	Check
	Raise
)

// Player: A structural representation of a physical Player
type Player struct {
	name  string
	chips float64
	cards Deck
}

// NewPlayer: Generate a new Player with chips taken either explicitly (via the chips argument) or implicitly (set in settings)
func NewPlayer(name string, chips float64, settings Settings) (*Player, error) {

	var err error
	if chips == 0 {
		if settings.StartingChips == 0 {
			err = errors.New("Cannot determine the amount of starting chips (was not given chips nor settings with set chips)")
		}
		chips = settings.StartingChips
	}
	return &Player{name: name, chips: chips}, err
}

// Player string representation
func (p *Player) String() string {
	return fmt.Sprintf("name: %s, chips: %f, cards: %s", p.name, p.chips, p.cards)
}

// Settings that help construct the table
type Settings struct {
	Ante          float64
	Dealer        Player
	SmallBlind    Player
	BigBlind      Player
	StartingChips float64
}
