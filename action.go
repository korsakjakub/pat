package main

import (
	"errors"
	"fmt"
)

type Action int

const (
	Fold Action = iota
	Check
	Raise
)

type Player struct {
	name  string
	chips float64
	cards Deck
}

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

func (p *Player) String() string {
	return fmt.Sprintf("name: %s, chips: %f, cards: %s", p.name, p.chips, p.cards)
}

type Settings struct {
	Ante          float64
	Dealer        Player
	SmallBlind    Player
	BigBlind      Player
	StartingChips float64
}
