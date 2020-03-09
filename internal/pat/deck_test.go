package pat

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
	drawnCards := deck.Draw(3)
	expectedCards := []Card{Card{Suit: D, Index: _Q},
		Card{Suit: D, Index: _K},
		Card{Suit: D, Index: _A},
	}
	if len(drawnCards.Cards) != 3 {
		t.Errorf("\n The lengths dont match (got %d, expected %d)\n", len(drawnCards.Cards), 3)
	} else {
		for i, card := range drawnCards.Cards {
			if card != expectedCards[i] {
				t.Errorf("\nGot %s,\nExpected %s", card, expectedCards[i])
			}
		}
	}
	t.Logf("new deck:\n%s\n", deck)
}

func TestNewCard(t *testing.T) {
	want := Card{Index: _J, Suit: S}
	got := NewCard(S, _J)
	if got.Suit != want.Suit || got.Index != want.Index {
		t.Errorf("didn't get the card. Got %s, expected %s\n", got, want)
	}
}
