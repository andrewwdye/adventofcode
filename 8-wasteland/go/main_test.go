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
		assert.Equal(t, 2, lo.Must(solve(strings.NewReader(input), false)))
	})

	t.Run("sample2", func(t *testing.T) {
		input := `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`
		assert.Equal(t, 6, lo.Must(solve(strings.NewReader(input), false)))
	})
}

func TestWalkGhosts(t *testing.T) {
	t.Run("sample1", func(t *testing.T) {
		input := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`
		assert.Equal(t, 6, lo.Must(solve(strings.NewReader(input), true)))
	})
}

func TestGcd(t *testing.T) {
	assert.Equal(t, 1, gcd(1, 1))
	assert.Equal(t, 6, gcd(48, 18))
}
