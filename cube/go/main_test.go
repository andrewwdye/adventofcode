package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestCountValid(t *testing.T) {
	tests := []struct {
		input    string
		r        int
		g        int
		b        int
		expected int
	}{
		{},
	}
	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, test.expected, lo.Must(countValid(strings.NewReader(test.input), test.r, test.g, test.b)))
		})
	}
}

func TestValidLine(t *testing.T) {
	assert.Equal(t, 5, lo.Must(validLine("Game 5: 1 red, 1 green, 1 blue", 1, 1, 1)))
	assert.Equal(t, 5, lo.Must(validLine("Game 5: 1 red, 1 green, 1 blue", 2, 2, 2)))
	assert.Equal(t, 0, lo.Must(validLine("Game 5: 1 red, 1 green, 1 blue", 0, 0, 0)))
}

func TestPower(t *testing.T) {
	assert.Equal(t, 48, lo.Must(power("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")))
	assert.Equal(t, 12, lo.Must(power("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue")))
	assert.Equal(t, 1560, lo.Must(power("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")))
	assert.Equal(t, 630, lo.Must(power("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red")))
	assert.Equal(t, 36, lo.Must(power("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green")))
}
