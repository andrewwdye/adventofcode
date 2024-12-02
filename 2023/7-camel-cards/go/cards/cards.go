package cards

import (
	"fmt"
	"slices"
)

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (t HandType) String() string {
	switch t {
	case HighCard:
		return "High Card"
	case OnePair:
		return "One Pair"
	case TwoPair:
		return "Two Pair"
	case ThreeOfAKind:
		return "Three of a Kind"
	case FullHouse:
		return "Full House"
	case FourOfAKind:
		return "Four of a Kind"
	case FiveOfAKind:
		return "Five of a Kind"
	default:
		return "Unknown"
	}
}

type Hand struct {
	Cards []Card
	Type  HandType
	Bet   int
}

type Set struct {
	Card  Card
	Count int
}

type Card int

const (
	Joker Card = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

func (c Card) String() string {
	switch c {
	case Joker:
		return "*"
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "10"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	default:
		return "Unknown"
	}
}

var cardMap = map[rune]Card{
	'2': Two,
	'3': Three,
	'4': Four,
	'5': Five,
	'6': Six,
	'7': Seven,
	'8': Eight,
	'9': Nine,
	'T': Ten,
	'J': Jack,
	'Q': Queen,
	'K': King,
	'A': Ace,
}

var cardMapJokers = map[rune]Card{
	'2': Two,
	'3': Three,
	'4': Four,
	'5': Five,
	'6': Six,
	'7': Seven,
	'8': Eight,
	'9': Nine,
	'T': Ten,
	'J': Joker,
	'Q': Queen,
	'K': King,
	'A': Ace,
}

func NewHand(handString string, bet int, jokers bool) Hand {
	cardMapper := cardMap
	if jokers {
		cardMapper = cardMapJokers
	}

	// Get cards
	cards := make([]Card, len(handString))
	for i, cardString := range handString {
		cards[i] = cardMapper[cardString]
	}

	// Get sets
	cardCounts := map[Card]int{}
	for _, card := range cards {
		cardCounts[card] += 1
	}
	sets := make([]Set, 0, len(cardCounts))
	for card, count := range cardCounts {
		if card == Joker {
			continue
		}
		sets = append(sets, Set{Card: card, Count: count})
	}
	slices.SortFunc(sets, func(a, b Set) int {
		if a.Count != b.Count {
			return b.Count - a.Count
		}
		return int(b.Card) - int(a.Card)
	})

	// Deal with jokers
	if len(sets) == 0 {
		sets = append(sets, Set{Card: Joker, Count: 0})
	}
	sets[0].Count += cardCounts[Joker]

	// Figure out hand type
	var handType HandType
	if sets[0].Count == 5 {
		handType = FiveOfAKind
	} else if sets[0].Count == 4 {
		handType = FourOfAKind
	} else if sets[0].Count == 3 && sets[1].Count == 2 {
		handType = FullHouse
	} else if sets[0].Count == 3 {
		handType = ThreeOfAKind
	} else if sets[0].Count == 2 && sets[1].Count == 2 {
		handType = TwoPair
	} else if sets[0].Count == 2 {
		handType = OnePair
	} else {
		handType = HighCard
	}

	return Hand{
		Cards: cards,
		Type:  handType,
		Bet:   bet,
	}
}

func (h Hand) Cmp(other Hand) int {
	if h.Type != other.Type {
		return int(h.Type) - int(other.Type)
	}
	for i := range h.Cards {
		if h.Cards[i] != other.Cards[i] {
			return int(h.Cards[i]) - int(other.Cards[i])
		}
	}
	return 0
}

func (h Hand) String() string {
	return fmt.Sprintf("%s --> %s: %d", h.Cards, h.Type, h.Bet)
}
