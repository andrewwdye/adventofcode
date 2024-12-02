package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHand(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		assert.Equal(t, FiveOfAKind, NewHand("AAAAA", 0, false).Type)
		assert.Equal(t, FourOfAKind, NewHand("AKKKK", 0, false).Type)
		assert.Equal(t, FullHouse, NewHand("AAKKK", 0, false).Type)
		assert.Equal(t, ThreeOfAKind, NewHand("AKQQQ", 0, false).Type)
		assert.Equal(t, TwoPair, NewHand("AKKQQ", 0, false).Type)
		assert.Equal(t, OnePair, NewHand("AKQJJ", 0, false).Type)
		assert.Equal(t, HighCard, NewHand("AKQJT", 0, false).Type)
	})

	t.Run("jokers", func(t *testing.T) {
		assert.Equal(t, FullHouse, NewHand("2233J", 0, true).Type)
	})
}
