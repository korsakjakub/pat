package pat

import (
	"testing"
)

func TestStraight(t *testing.T) {
	deck := NewDeck()
	player := Deck{Cards: []Card{deck.Cards[18], deck.Cards[19]}}
	table := Deck{Cards: []Card{deck.Cards[20], deck.Cards[21], deck.Cards[22]}}

	if GetHand(player, table) != Straight {
		t.Errorf("didn't get the expected Hand!. Got %s, expected %s\n", GetHand(player, table), Straight)
	}
}

func TestWinningHand(t *testing.T) {
	deck := NewDeck()
	a := Deck{Cards: []Card{deck.Cards[21], deck.Cards[22], deck.Cards[23], deck.Cards[24], deck.Cards[25]}}

	b := Deck{Cards: []Card{deck.Cards[18], deck.Cards[19], deck.Cards[20], deck.Cards[21], deck.Cards[22]}}

	t.Log(WinningHand(a, b))
}
