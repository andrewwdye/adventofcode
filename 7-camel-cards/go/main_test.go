package main

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestSumHands(t *testing.T) {
	t.Run("sample", func(t *testing.T) {
		input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
		assert.Equal(t, 6440, lo.Must(sumHands(strings.NewReader(input))))
	})
}

func TestNewHand(t *testing.T) {
	assert.Equal(t, FiveOfAKind, NewHand("AAAAA", 0).Type)
	assert.Equal(t, FourOfAKind, NewHand("AKKKK", 0).Type)
	assert.Equal(t, FullHouse, NewHand("AAKKK", 0).Type)
	assert.Equal(t, ThreeOfAKind, NewHand("AKQQQ", 0).Type)
	assert.Equal(t, TwoPair, NewHand("AKKQQ", 0).Type)
	assert.Equal(t, OnePair, NewHand("AKQJJ", 0).Type)
	assert.Equal(t, HighCard, NewHand("AKQJT", 0).Type)
}
