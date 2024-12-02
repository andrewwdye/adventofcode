package pkg

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestCheckColumn(t *testing.T) {
	input1 := strings.Split(`#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.`, "\n")
	input2 := strings.Split(`#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`, "\n")

	assert.Greater(t, checkColumn(input1, 1), 0)
	assert.Greater(t, checkColumn(input1, 2), 0)
	assert.Equal(t, checkColumn(input1, 5), 0)

	for i := 1; i < len(input2[0])-1; i++ {
		assert.Greater(t, checkColumn(input2, i), 0)
	}
}

func TestCheckRow(t *testing.T) {
	input1 := strings.Split(`#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.`, "\n")
	input2 := strings.Split(`#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`, "\n")

	for i := 1; i < len(input1)-1; i++ {
		assert.Greater(t, checkRow(input1, i), 0)
	}

	assert.Greater(t, checkRow(input2, 3), 0)
	assert.Equal(t, checkRow(input2, 4), 0)
}

func TestFindReflection(t *testing.T) {
	input1 := strings.Split(`#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.`, "\n")
	input2 := strings.Split(`#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`, "\n")

	input3 := strings.Split(`#.####.#...
#.##.#.#...
###.#.####.
.##..#####.
.#.#.#.#..#
..#......#.
..#......#.`, "\n")

	col, row := findReflection(input1, 0)
	assert.Equal(t, 5, col)
	assert.Equal(t, 0, row)

	col, row = findReflection(input2, 0)
	assert.Equal(t, 0, col)
	assert.Equal(t, 4, row)

	col, row = findReflection(input3, 0)
	assert.Equal(t, 0, col)
	assert.Equal(t, 6, row)
}

func TestSolve(t *testing.T) {
	input := `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

	assert.Equal(t, 405, lo.Must(Solve(strings.NewReader(input), 0)))
}

func TestSolve2(t *testing.T) {
	input := `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

	assert.Equal(t, 400, lo.Must(Solve(strings.NewReader(input), 1)))
}
