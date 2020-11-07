package pat

import (
	"fmt"
	"testing"
)

func TestStrings(t *testing.T) {
	categories := []category{Straight, Pair, TwoPair}
	want := []string{"Straight", "Pair", "Two pairs"}
	for i, cat := range categories {
		got := fmt.Sprintf("%s", cat)
		if got != want[i] {
			t.Errorf("Wanted %s, got %s\n", want[i], got)
		}
	}
}

func TestGetHand(t *testing.T) {
	var testValues = []struct {
		player Deck // input from player
		table  Deck // input from table
		want   Hand // output
	}{
		{
			Deck([]Card{NewCard(S, _2), NewCard(D, _3)}),
			Deck([]Card{NewCard(S, _4), NewCard(D, _5), NewCard(S, _6)}),
			Hand{category: Straight, high: NewCard(S, _6)},
		},
		{
			Deck([]Card{NewCard(S, _2), NewCard(S, _4)}),
			Deck([]Card{NewCard(S, _J), NewCard(S, _5), NewCard(S, _6)}),
			Hand{category: Flush, high: NewCard(S, _J)},
		},
		{
			Deck([]Card{NewCard(S, _2), NewCard(S, _3)}),
			Deck([]Card{NewCard(S, _4), NewCard(S, _5), NewCard(S, _6)}),
			Hand{category: RoyalFlush, high: NewCard(S, _6)},
		},
		{
			Deck([]Card{NewCard(D, _2), NewCard(D, _3)}),
			Deck([]Card{NewCard(D, _4), NewCard(D, _5), NewCard(D, _6)}),
			Hand{category: StraightFlush, high: NewCard(D, _6)},
		},
		{ // unordered, but straight
			Deck([]Card{NewCard(S, _2), NewCard(C, _4)}),
			Deck([]Card{NewCard(D, _6), NewCard(D, _3), NewCard(D, _5)}),
			Hand{category: Straight, high: NewCard(D, _6)},
		},
		{
			Deck([]Card{NewCard(S, _2), NewCard(C, _2)}),
			Deck([]Card{NewCard(D, _2), NewCard(H, _2), NewCard(D, _5)}),
			Hand{category: Quads, high: NewCard(S, _2)},
		},
		{
			Deck([]Card{NewCard(S, _2), NewCard(C, _2)}),
			Deck([]Card{NewCard(D, _3), NewCard(H, _2), NewCard(D, _5)}),
			Hand{category: Trips, high: NewCard(S, _2)},
		},
		{
			Deck([]Card{NewCard(S, _2), NewCard(C, _2)}),
			Deck([]Card{NewCard(D, _3), NewCard(H, _2), NewCard(H, _3)}),
			Hand{category: FullHouse, high: NewCard(S, _2), sHigh: NewCard(H, _3)},
		},
		{
			Deck([]Card{NewCard(S, _2), NewCard(C, _2)}),
			Deck([]Card{NewCard(D, _5), NewCard(H, _7), NewCard(H, _J)}),
			Hand{category: Pair, high: NewCard(S, _2)},
		},
		{
			Deck([]Card{NewCard(S, _2), NewCard(C, _2)}),
			Deck([]Card{NewCard(D, _3), NewCard(H, _K), NewCard(H, _3)}),
			Hand{category: TwoPair, high: NewCard(S, _3), sHigh: NewCard(S, _2)},
		},
		{
			Deck([]Card{NewCard(S, _2), NewCard(C, _9)}),
			Deck([]Card{NewCard(D, _3), NewCard(H, _K), NewCard(H, _A)}),
			Hand{category: High, high: NewCard(H, _A)},
		},
	}

	for _, el := range testValues {
		if check := GetHand(el.player, el.table); check.category != el.want.category || check.high != el.want.high {
			t.Errorf("didn't get the expected Hand!. Got %s %s, expected %s %s\n", check.category, check.high, el.want.category, el.want.high)
		}
	}
}

func TestWinningHand(t *testing.T) {
	//a := Deck([]Card{NewCard(S, _2), NewCard(D, _3), NewCard(D, _2), NewCard(C, _J), NewCard(H, _A)})
	//b := Deck([]Card{NewCard(S, _3), NewCard(H, _3), NewCard(D, _3), NewCard(C, _J), NewCard(H, _A)})

	//t.Log(WinningHand(a, b))
}

func TestMapCardOccurrences(t *testing.T) {
	var testValues = []struct {
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

	for _, e := range testValues {
		got := mapCards(e.cards)
		if got[e.position] != e.amount {
			t.Errorf("didnt get the map i wanted. Got %v, wanted %d\n", got[e.position], e.amount)
		}
	}
}

func TestIsBetterThan(t *testing.T) {
	var testValues = []struct {
		c1       Hand
		c2       Hand
		isBetter Equivalence
	}{
		{
			c1:       Hand{category: Quads, high: NewCard(S, _6)},
			c2:       Hand{category: Quads, high: NewCard(S, _5)},
			isBetter: better,
		},
		{
			c1:       Hand{category: Quads, high: NewCard(S, _5)},
			c2:       Hand{category: Quads, high: NewCard(S, _6)},
			isBetter: worse,
		},
		{
			c1:       Hand{category: Quads, high: NewCard(S, _5)},
			c2:       Hand{category: Trips, high: NewCard(S, _6)},
			isBetter: better,
		},
		{
			c1:       Hand{category: FullHouse, high: NewCard(S, _6), sHigh: NewCard(S, _5)},
			c2:       Hand{category: FullHouse, high: NewCard(S, _6), sHigh: NewCard(S, _2)},
			isBetter: better,
		},
		{
			c1:       Hand{category: FullHouse, high: NewCard(S, _6), sHigh: NewCard(S, _2)},
			c2:       Hand{category: FullHouse, high: NewCard(S, _6), sHigh: NewCard(S, _5)},
			isBetter: worse,
		},
		{
			c1:       Hand{category: FullHouse, high: NewCard(S, _6), sHigh: NewCard(S, _2)},
			c2:       Hand{category: FullHouse, high: NewCard(S, _6), sHigh: NewCard(S, _2)},
			isBetter: equal,
		},
	}

	for _, e := range testValues {
		if e.c1.IsBetterThan(e.c2) != e.isBetter {
			t.Errorf("Got %v, wanted %v\n", e.c1.IsBetterThan(e.c2), e.isBetter)
		}
	}
}
