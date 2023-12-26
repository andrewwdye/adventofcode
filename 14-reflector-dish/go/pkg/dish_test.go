package pkg

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTilt(t *testing.T) {
	input := `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

	assert.Equal(t, 136, tilt(strings.Split(input, "\n")))
}
