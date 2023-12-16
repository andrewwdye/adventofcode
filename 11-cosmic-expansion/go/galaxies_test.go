package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseUniverse(t *testing.T) {
	input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
	univ := ParseUniverse(strings.NewReader(input))
	assert.Len(t, univ.Galaxies, 9, univ)
	assert.Equal(t, 13, univ.Width)
	assert.Equal(t, 12, univ.Height)
}
