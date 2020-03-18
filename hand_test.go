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
	a := Deck([]Card{NewCard(S, _2), NewCard(D, _3), NewCard(D, _2), NewCard(C, _J), NewCard(H, _A)})
	a := Deck([]Card{NewCard(S, _3), NewCard(H, _3), NewCard(D, _3), NewCard(C, _J), NewCard(H, _A)})

	t.Log(WinningHand(a, b))
}

func TestMapCardOccurrences(t *testing.T) {
	var test_values = []struct {
		cards    Deck // input from player
		position Index
		amount   int // output
	}{
		{
			Deck([]Card{NewCard(S, _2), NewCard(D, _2), NewCard(C, _2), NewCard(H, _2), NewCard(H, _4)}),
			_2,
			4,
		},
		{
			Deck([]Card{NewCard(S, _K), NewCard(D, _2), NewCard(C, _K), NewCard(H, _K), NewCard(H, _5)}),
			_K,
			3,
		},
		{
			Deck([]Card{NewCard(S, _J), NewCard(D, _2), NewCard(C, _J), NewCard(H, _3), NewCard(H, _5)}),
			_J,
			2,
		},
		{ // next 2 check full house
			Deck([]Card{NewCard(S, _J), NewCard(D, _3), NewCard(C, _J), NewCard(H, _3), NewCard(S, _3)}),
			_J,
			2,
		},
		{
			Deck([]Card{NewCard(S, _J), NewCard(D, _3), NewCard(C, _J), NewCard(H, _3), NewCard(S, _3)}),
			_3,
			3,
		},
	}

	for _, e := range test_values {
		got := mapCardOccurrences(e.cards)
		if got[e.position] != e.amount {
			t.Errorf("didnt get the map i wanted. Got %v, wanted %d\n", got[e.position], e.amount)
		}
	}
}
