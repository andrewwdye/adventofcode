package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func sumHandsFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return sumHands(file)
}

func sumHands(reader io.Reader) (int, error) {
	hands, err := getHands(reader)
	if err != nil {
		return 0, err
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		return a.Cmp(b)
	})

	sum := 0
	for i, hand := range hands {
		// fmt.Printf("%04d: %s --> %d\n", i+1, hand, hand.Bet*(i+1))
		sum += hand.Bet * (i + 1)
	}

	fmt.Println(sum)
	return sum, nil
}

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
	Cards        string
	Type         HandType
	OrderedCards []Card
	Bet          int
}

type Card int

const (
	Two Card = iota
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

func NewHand(cards string, bet int) Hand {
	counts := map[Card]int{}
	for _, c := range cards {
		counts[cardMap[c]] += 1
	}
	groups := map[int][]Card{
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
	}
	for card, count := range counts {
		groups[count] = append(groups[count], card)
	}
	for _, group := range groups {
		slices.SortFunc(group, func(a, b Card) int {
			return int(b) - int(a)
		})
	}
	ordered := make([]Card, 0, len(cards))
	for _, card := range cards {
		ordered = append(ordered, cardMap[card])
	}
	hand := Hand{
		Cards:        cards,
		Bet:          bet,
		OrderedCards: ordered,
	}
	if len(groups[5]) > 0 {
		hand.Type = FiveOfAKind
	} else if len(groups[4]) > 0 {
		hand.Type = FourOfAKind
	} else if len(groups[3]) > 0 && len(groups[2]) > 0 {
		hand.Type = FullHouse
	} else if len(groups[3]) > 0 {
		hand.Type = ThreeOfAKind
	} else if len(groups[2]) > 1 {
		hand.Type = TwoPair
	} else if len(groups[2]) > 0 {
		hand.Type = OnePair
	} else {
		hand.Type = HighCard
	}
	return hand
}

func (h Hand) Cmp(other Hand) int {
	if h.Type != other.Type {
		return int(h.Type) - int(other.Type)
	}
	for i := range h.OrderedCards {
		if h.OrderedCards[i] != other.OrderedCards[i] {
			return int(h.OrderedCards[i]) - int(other.OrderedCards[i])
		}
	}
	return 0
}

func (h Hand) String() string {
	return fmt.Sprintf("%s --> %s: %d", h.Cards, h.Type, h.Bet)
}

func getHands(reader io.Reader) ([]Hand, error) {
	scanner := bufio.NewScanner(reader)
	re := regexp.MustCompile(`([0-9AKQJT]{5}) (\d+)`)
	hands := []Hand{}
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) != 3 {
			return nil, fmt.Errorf("invalid input: %s", line)
		}
		hand := matches[1]
		bet, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, fmt.Errorf("invalid input: %s", line)
		}
		hands = append(hands, NewHand(hand, bet))
	}
	return hands, nil
}

func main() {
	if _, err := sumHandsFile(os.Args[1]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
