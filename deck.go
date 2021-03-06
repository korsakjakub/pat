package pat

import (
	"fmt"
	"math/rand"
	"time"
)

// Suit - Possible suits of standard playing cards
type Suit int

// Spades, Hearts, Clubs, Diamonds
const (
	S Suit = iota + 1 // Spade
	H                 // Heart
	C                 // Club
	D                 // Diamond
)

// Index - Possible indices of standard playing cards
type Index int

// All card faces from 2, to an Ace
const (
	_2 Index = iota + 2
	_3
	_4
	_5
	_6
	_7
	_8
	_9
	_10
	_J
	_Q
	_K
	_A
)

func max(indices []Index) (out Index) {
	for _, i := range indices {
		if i > out {
			out = i
		}
	}
	return
}

// SuitEnumerate gives all suits as a slice
func SuitEnumerate() []Suit {
	var res []Suit
	for i := S; i <= D; i++ {
		res = append(res, i)
	}
	return res
}

// Suit string representation. Keep in mind the s-1, for intuition we start counting suits from 1
func (s Suit) String() string {
	return [...]string{"S", "H", "C", "D"}[s-1]
}

// Index string representation. Keep in mind the i-1, for intuition we start counting indices from 2
func (i Index) String() string {
	return [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}[i-2]
}

// IndexEnumerate gives all indices as a slice
func IndexEnumerate() []Index {
	var res []Index
	for i := _2; i <= _A; i++ {
		res = append(res, i)
	}
	return res
}

// Card - A playing Card is clearly mapped by a pair (Suit, Index)
type Card struct {
	Suit  Suit
	Index Index
}

// NewCard returns a card when given a suit and an index.
func NewCard(suit Suit, index Index) Card {
	return Card{Suit: suit, Index: index}
}

// A Deck is a slice of Cards
type Deck []Card

// Deck string representation. Each line will print 13 cards (if there are so many)
func (d Deck) String() string {
	res := ""
	for i, card := range d {
		res += fmt.Sprintf("%s%s", card.Index, card.Suit)
		if i%13 != 12 {
			res += "\t"
		} else {
			res += "\n"
		}
	}
	return res
}

// NewDeck generates a new Deck with standard order [low -> high] [spades -> hearts -> clubs -> diamonds]
func NewDeck() Deck {
	cards := Deck{}
	for _, suit := range SuitEnumerate() {
		for _, index := range IndexEnumerate() {
			cards = append(cards, Card{Suit: suit, Index: index})
		}
	}
	return Deck(cards)
}

// Rm (remove) a Card from Deck. The resulting Deck should consist of n-1 cards, if the input Deck had n (any Deck can only be a permutation of a standard Deck).
func (d Deck) Rm(c Card) {
	for i, card := range d {
		if card.Suit == c.Suit && card.Index == c.Index {
			d = append(d[:i], d[i+1:]...)
		}
	}
}

// Draw n cards from the top. The resulting Deck should have L - n cards, where L denotes the amount of cards before the draw.
func (d Deck) Draw(n int) Deck {
	// it's important to first draw the cards, then remove them, not in the opposite order
	drawnCards := Deck(d[len(d)-n:])
	d = d[:len(d)-n]
	return drawnCards
}

// Shuffle the Deck. Note: it should always generate a permutation of the input Deck.
func (d Deck) Shuffle(s ...int64) {
	seed := int64(0)
	if len(s) == 0 {
		seed = time.Now().UnixNano()
	} else if len(s) >= 2 {
		for _, el := range s {
			seed += el
		}
		seed /= 2
	} else {
		seed = s[0]
	}
	rand.Seed(seed)
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}

func (d Deck) Max() Index {
	var maxIndex Index
	for i, e := range d {
		if i == 0 || e.Index >= maxIndex {
			maxIndex = e.Index
		}
	}
	return maxIndex
}

func (d Deck) Len() int { return len(d) }

func (d Deck) Less(a, b int) bool { return d[a].Index < d[b].Index }

func (d Deck) Swap(a, b int) { d[a], d[b] = d[b], d[a] }
