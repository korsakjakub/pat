package main

import (
	"testing"
	//"time"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	if len(deck.Cards) != 52 {
		t.Error("Incorrect amount of cards in a new deck!")
	}
	t.Log(deck)
}

func TestRemoveCard(t *testing.T) {
	deck := NewDeck()
	card := Card{Suit: S, Index: _A}
	deck.Rm(card)

	for _, d := range deck.Cards {
		if d.Suit == S && d.Index == _A {
			t.Error("Card is still present in the deck")
		}
	}
	t.Log(deck)
}

func TestShuffle(t *testing.T) {
	deck := NewDeck()
	//deck.Shuffle(time.Now().UnixNano())
	deck.Shuffle(1)
	//t.Error(deck)
}

func TestDraw(t *testing.T) {
	deck := NewDeck()
	drawn_cards := deck.Draw(3)
	ex_cards := []Card{Card{Suit: D, Index: _Q},
		Card{Suit: D, Index: _K},
		Card{Suit: D, Index: _A},
	}
	if len(drawn_cards.Cards) != 3 {
		t.Errorf("\n The lenghts dont match (got %d, expected %d)\n", len(drawn_cards.Cards), 3)
	} else {
		for i, card := range drawn_cards.Cards {
			if card != ex_cards[i] {
				t.Errorf("\nGot %s,\nExpected %s", card, ex_cards[i])
			}
		}
	}
	t.Logf("new deck:\n%s\n", deck)
}
