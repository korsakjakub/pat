package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Suit int

const (
	S Suit = iota + 1 // Spade
	H                 // Heart
	C                 // Club
	D                 // Diamond
)

const (
	_2 Index = iota + 1
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

func SuitEnumerate() []Suit {
	var res []Suit
	for i := S; i <= D; i++ {
		res = append(res, i)
	}
	return res
}

func (s Suit) String() string {
	return [...]string{"S", "H", "C", "D"}[s-1]
}

type Index int

func (i Index) String() string {
	return [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}[i-1]
}

func IndexEnumerate() []Index {
	var res []Index
	for i := _2; i <= _A; i++ {
		res = append(res, i)
	}
	return res
}

type Card struct {
	Suit  Suit
	Index Index
}

type Deck struct {
	Cards []Card
}

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

func NewDeck() *Deck {
	cards := []Card{}
	for _, suit := range SuitEnumerate() {
		for _, index := range IndexEnumerate() {
			cards = append(cards, Card{Suit: suit, Index: index})
		}
	}
	return &Deck{cards}
}

func (d *Deck) Rm(c Card) {
	for i, card := range d.Cards {
		if card.Suit == c.Suit && card.Index == c.Index {
			d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
		}
	}
}

func (d *Deck) Draw(n int) Deck {
	// it's important to first draw the cards, then remove them, not in the oposite order
	drawn_cards := Deck{Cards: d.Cards[52-n:]}
	d.Cards = d.Cards[:len(d.Cards)-n]
	return drawn_cards
}

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
