package pat

import (
	"sort"
)

// category - Possible Hands Categories that a Player can get
type category int

// All the hand categories in standard poker
const (
	High category = iota
	Pair
	TwoPair
	Trips
	Straight
	Flush
	FullHouse
	Quads
	StraightFlush
	RoyalFlush
)

type Equivalence int

const (
	worse Equivalence = iota
	better
	equal
)

func (e Equivalence) String() string {
	return []string{"worse", "better", "equal"}[e]
}

// Hand is just a category with a rank. We first compare Categories, then ranks.
type Hand struct {
	category
	high  Card
	sHigh Card
}

// IsBetterThan decides i
func (h Hand) IsBetterThan(o Hand) Equivalence {
	if h.category == o.category && h.high.Index == o.high.Index && h.sHigh.Index == o.sHigh.Index {
		return equal
	}
	if h.category < o.category {
		return worse
	} else if h.category == o.category {
		if h.high.Index < o.high.Index {
			return worse
		}
		if h.sHigh.Index < o.sHigh.Index {
			return worse
		}
	}
	return better
}

// category String representation
func (c category) String() string {
	return []string{"High card", "Pair", "Two pairs", "Trips", "Straight", "Flush", "Full House", "Quads", "Straight Flush", "Royal Flush"}[c]
}

func isFlush(cards Deck) bool {
	suit := cards[0].Suit
	for _, card := range cards {
		if card.Suit != suit {
			return false
		}
	}
	return true
}

func isStraight(cards Deck) (bool, Card) {
	index := cards[0].Index
	for _, card := range cards {
		if index != card.Index {
			return false, Card{}
		}
		index++
	}
	return true, cards[len(cards)-1]
}

type occurrences map[Index]int

func (o occurrences) findIndex(i int) []Index {
	var out []Index
	for k := range o {
		if o[k] == i {
			out = append(out, k)
		}
	}
	return out
}

func mapCards(cards Deck) occurrences {
	o := make(occurrences)
	for _, ci := range cards {
		o[ci.Index]++
	}
	return o
}

func checkHands(cards Deck) Hand {
	sort.Sort(cards)

	// check flush, straight flush and royal flush. Note: a flush discriminates any pairs, trips or quads!
	if isFlush(cards) {
		if isIt, high := isStraight(cards); isIt {
			if cards[0].Suit == S {
				return Hand{category: RoyalFlush, high: high}
			}
			return Hand{category: StraightFlush, high: cards[len(cards)-1]}
		}
		return Hand{category: Flush, high: cards[len(cards)-1]}
	}

	// there only can be a straight that is not a Royal (or a straight) Flush
	if isIt, high := isStraight(cards); isIt {
		return Hand{category: Straight, high: high}
	}

	o := mapCards(cards)

	trips := []Index{}
	pairs := []Index{}

	for index := range o {
		switch o[index] {
		case 4:
			return Hand{category: Quads, high: NewCard(S, index)}
		case 3:
			trips = append(trips, index)
		case 2:
			pairs = append(pairs, index)
		}
	}

	if len(trips) > 0 {
		switch len(pairs) {
		case 0:
			return Hand{category: Trips, high: NewCard(S, trips[0])}
		case 1:
			return Hand{category: FullHouse, high: NewCard(S, max(trips)), sHigh: NewCard(S, max(pairs))}
		}
	} else {
		switch len(pairs) {
		case 1:
			return Hand{category: Pair, high: NewCard(S, pairs[0])}
		case 2:
			return Hand{category: TwoPair, high: NewCard(S, max(pairs))}
		}
	}
	return Hand{category: High, high: cards[len(cards)-1]}
}

func getBestFive(cards Deck) Deck {
	if len(cards) <= 5 {
		return cards
	}
	return cards
}

// GetHand returns a hand provided cards from a player and the table
func GetHand(player, table Deck) Hand {
	cards := Deck{}
	switch phase := len(table); phase {
	case 0: // preflop
		cards = player
	case 3: // flop
		cards = append(append(cards, player...), table...)
	case 4: // turn
	case 5: // river
		cards = getBestFive(append(player, table...))
	}
	return checkHands(cards)
}
