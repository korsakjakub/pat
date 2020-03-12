package pat

import (
	"testing"
)

func TestGetHand(t *testing.T) {
	var test_values = []struct {
		player Deck     // input from player
		table  Deck     // input from table
		want   Category // output
	}{
		{
			Deck([]Card{NewCard(S, _2), NewCard(D, _3)}),
			Deck([]Card{NewCard(S, _4), NewCard(D, _5), NewCard(S, _6)}),
			Straight,
		},
		{
			Deck([]Card{NewCard(S, _2), NewCard(S, _4)}),
			Deck([]Card{NewCard(S, _J), NewCard(S, _5), NewCard(S, _6)}),
			Flush,
		},
		{
			Deck([]Card{NewCard(S, _2), NewCard(S, _3)}),
			Deck([]Card{NewCard(S, _4), NewCard(S, _5), NewCard(S, _6)}),
			RoyalFlush,
		},
		{
			Deck([]Card{NewCard(D, _2), NewCard(D, _3)}),
			Deck([]Card{NewCard(D, _4), NewCard(D, _5), NewCard(D, _6)}),
			StraightFlush,
		},
		{ // unordered, but straight
			Deck([]Card{NewCard(S, _2), NewCard(C, _4)}),
			Deck([]Card{NewCard(D, _6), NewCard(D, _3), NewCard(D, _5)}),
			Straight,
		},
	}

	for _, el := range test_values {
		if check := GetHand(el.player, el.table); check.Category != el.want {
			t.Errorf("didn't get the expected Hand!. Got %s, expected %s\n", check, el.want)
		}
	}
}

func TestWinningHand(t *testing.T) {
	deck := NewDeck()
	a := Deck([]Card{deck[21], deck[22], deck[23], deck[24], deck[25]})

	b := Deck([]Card{deck[18], deck[19], deck[20], deck[21], deck[22]})

	t.Log(WinningHand(a, b))
}

func TestMapCardOccurrences(t *testing.T) {
	cards := []Card{NewCard(S, _2), NewCard(D, _2), NewCard(C, _2), NewCard(H, _2), Card{}}

	got := mapCardOccurrences(cards)
	if got[_2] != 4 {
		t.Errorf("didnt get the map i wanted. Got %d, wanted %s\n", got[_2], "4")
	}
}
