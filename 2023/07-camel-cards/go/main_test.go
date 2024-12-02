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
		assert.Equal(t, 5905, lo.Must(sumHands(strings.NewReader(input))))
		//6440
	})
}
