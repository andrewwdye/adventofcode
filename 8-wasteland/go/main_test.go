package main

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	t.Run("sample1", func(t *testing.T) {
		input := `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`
		assert.Equal(t, 2, lo.Must(solve(strings.NewReader(input))))
	})

	t.Run("sample2", func(t *testing.T) {
		input := `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`
		assert.Equal(t, 6, lo.Must(solve(strings.NewReader(input))))
	})
}
