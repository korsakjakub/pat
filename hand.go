package pat

import (
	"fmt"
	"sort"
	"strconv"
)

// Category - Possible Hands Categories that a Player can get
type Category int

// All the hand categories in standard poker
const (
	High Category = iota
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

// Hand is just a Category with a Rank. We first compare Categories, then Ranks.
type Hand struct {
	Category Category
	Rank     int
}

// Category String representation
func (c Category) String() string {
	return []string{"High card", "Pair", "Two pairs", "Trips", "Straight", "Flush", "Full House", "Quads", "Straight Flush", "Royal Flush"}[c]
}

func (h Hand) String() string {
	return fmt.Sprintf("%s, %s", h.Category, strconv.Itoa(h.Rank))
}

func isFlush(cards []Card) bool {
	suit := cards[0].Suit
	for _, card := range cards {
		if card.Suit != suit {
			return false
		}
	}
	return true
}

func isStraight(cards []Card) bool {
	index := cards[0].Index
	for _, card := range cards {
		if index != card.Index {
			return false
		}
		index++
	}
	return true
}

func mapCardOccurrences(cards []Card) map[Index]int {
	occurrences := make(map[Index]int)
	for _, ci := range cards {
		occurrences[ci.Index]++
	}
	return occurrences
}

func checkHands(cards Deck) Hand {
	sort.Sort(cards)

	// check flush, straight flush and royal flush. Note: a flush discriminates any pairs, trips or quads!
	if isFlush(cards) {
		if isStraight(cards) {
			if cards[0].Suit == S {
				return Hand{Category: RoyalFlush}
			}
			return Hand{Category: StraightFlush}
		}
		return Hand{Category: Flush}
	}

	// there only can be a straight that is not a Royal (or a straight) Flush
	if isStraight(cards) {
		return Hand{Category: Straight}
	}

	occurrences := mapCardOccurrences(cards)
	switch occurrences {

	}

	// if there are quads
	if len(occurrences) == 2 {
		return Hand{Category: Quads}
	}

	/*
		// if there is a full house
		if isFullHouse(cards) {
			return FullHouse
		}

		// if there are trips
		if isTrips(cards) {
			return Trips
		}

		// if there are two pairs
		if isTwoPair(cards) {
			return TwoPair
		}

		// if there is a pair
		if isPair(cards) {
			return Pair
		}
	*/
	return Hand{Category: High}
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

// WinningHand returns the winning hand
func WinningHand(a, b Deck) *Deck {
	return &Deck{}
}
