package pkg // import github.com/korsakjakub/pat/pkg

import (
	"fmt"
	"math/rand"
	"time"
)

// Possible suits of standard playing cards
type Suit int

const (
	S Suit = iota + 1 // Spade
	H                 // Heart
	C                 // Club
	D                 // Diamond
)

// Possible indices of standard playing cards
type Index int

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

// Gives all suits as a slice
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

// Gives all indices as a slice
func IndexEnumerate() []Index {
	var res []Index
	for i := _2; i <= _A; i++ {
		res = append(res, i)
	}
	return res
}

// A playing Card is clearly mapped by a pair (Suit, Index)
type Card struct {
	Suit  Suit
	Index Index
}

// A Deck is a slice of Cards
type Deck struct {
	Cards []Card
}

// Deck string representation. Each line will print 13 cards (if there are so many)
func (d Deck) String() string {
	res := ""
	for i, card := range d.Cards {
		res += fmt.Sprintf("%s%s", card.Index, card.Suit)
		if i%13 != 12 {
			res += "\t"
		} else {
			res += "\n"
		}
	}
	return res
}

// Generate a new Deck with standard order [low -> high] [spades -> hearts -> clubs -> diamonds]
func NewDeck() *Deck {
	cards := []Card{}
	for _, suit := range SuitEnumerate() {
		for _, index := range IndexEnumerate() {
			cards = append(cards, Card{Suit: suit, Index: index})
		}
	}
	return &Deck{cards}
}

// Remove a Card from Deck. The resulting Deck should consist of n-1 cards, if the input Deck had n (any Deck can only be a permutation of a standard Deck).
func (d *Deck) Rm(c Card) {
	for i, card := range d.Cards {
		if card.Suit == c.Suit && card.Index == c.Index {
			d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
		}
	}
}

// Draw n cards from the top. The resulting Deck should have L - n cards, where L denotes the amount of cards before the draw.
func (d *Deck) Draw(n int) Deck {
	// it's important to first draw the cards, then remove them, not in the oposite order
	drawn_cards := Deck{Cards: d.Cards[len(d.Cards)-n:]}
	d.Cards = d.Cards[:len(d.Cards)-n]
	return drawn_cards
}

// Shuffle the Deck. Note: it should always generate a permutation of the input Deck.
func (d *Deck) Shuffle(s ...int64) {
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
	c := d.Cards
	rand.Seed(seed)
	rand.Shuffle(len(c), func(i, j int) { c[i], c[j] = c[j], c[i] })
}
