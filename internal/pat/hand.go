package pat

// Hand - Possible Hands that a Player can get
type Hand int

const (
	High Hand = iota
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

func (h Hand) String() string {
	return []string{"High card", "Pair", "Two pairs", "Trips", "Straight", "Flush", "Full House", "Quads", "Straight Flush", "Royal Flush"}[h]
}

func checkHands(cards []Card) Hand {
	return High
}

func getBestFive(cards []Card) []Card {
	return []Card{}
}

func GetHand(player, table Deck) Hand {
	cards := []Card{}
	switch phase := len(table.Cards); phase {
	case 0: // preflop
		cards = player.Cards
	case 3: // flop
		cards = append(append(cards, player.Cards...), table.Cards...)
	case 4: // turn
	case 5: // river
		cards = getBestFive(append(player.Cards, table.Cards...))
	}
	return checkHands(cards)
}

func WinningHand(a Deck, b Deck) *Deck {
	return &Deck{}
}
