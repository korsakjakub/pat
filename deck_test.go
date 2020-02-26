package main

import (
	"testing"
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
	card := Card{Suit: S, Index: "A"}
	deck.Rm(card)

	for _, d := range deck.Cards {
		if d.Suit == S && d.Index == "A" {
			t.Error("Card is still present in the deck")
		}
	}
	t.Log(deck)
}
