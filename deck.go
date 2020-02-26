package main

import "fmt"

type Suit int

const (
	S Suit = iota + 1 // Spade
	H                 // Heart
	C                 // Club
	D                 // Diamond
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

type Index string

func IndexEnumerate() []Index {
	return []Index{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
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
		if i%13 != 12 {
			res += fmt.Sprintf("%s%s\t", card.Index, card.Suit)
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

}
