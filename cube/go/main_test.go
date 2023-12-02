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
