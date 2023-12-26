package pkg

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestTiltUp(t *testing.T) {
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

	expected := `OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....`

	result := tiltUp(strings.Split(input, "\n"))
	assert.Equal(t, strings.Split(expected, "\n"), result, strings.Join(result, "\n"))
}

func TestSolve1(t *testing.T) {
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

	assert.Equal(t, 136, lo.Must(Solve1(strings.NewReader(input))))
}
