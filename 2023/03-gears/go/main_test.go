package main

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestIsLocSymbol(t *testing.T) {
	t.Run("symbols", func(t *testing.T) {
		assert.True(t, isLocSymbol(0, 0, []string{"@"}))
		assert.False(t, isLocSymbol(0, 0, []string{"0"}))
		assert.False(t, isLocSymbol(0, 0, []string{"."}))
	})

	t.Run("out of bounds", func(t *testing.T) {
		assert.False(t, isLocSymbol(0, -1, []string{"@"}))
		assert.False(t, isLocSymbol(0, 1, []string{"@"}))
		assert.False(t, isLocSymbol(-1, 0, []string{"@"}))
		assert.False(t, isLocSymbol(1, 0, []string{"@"}))
	})
}

func TestIsPart(t *testing.T) {
	t.Run("above", func(t *testing.T) {
		assert.NotEqual(t, uint8(0), getPart([]int{1, 2}, 1, []string{
			"@..",
			".0.",
			"...",
		}).symbol)
		assert.NotEqual(t, uint8(0), getPart([]int{1, 2}, 1, []string{
			".@.",
			".0.",
			"...",
		}).symbol)
		assert.NotEqual(t, uint8(0), getPart([]int{1, 2}, 1, []string{
			"..@",
			".0.",
			"...",
		}).symbol)
	})

	t.Run("below", func(t *testing.T) {
		assert.NotEqual(t, uint8(0), getPart([]int{1, 2}, 1, []string{
			"...",
			".0.",
			"@..",
		}).symbol)
		assert.NotEqual(t, uint8(0), getPart([]int{1, 2}, 1, []string{
			"...",
			".0.",
			".@.",
		}).symbol)
		assert.NotEqual(t, uint8(0), getPart([]int{1, 2}, 1, []string{
			"...",
			".0.",
			"..@",
		}).symbol)
	})

	t.Run("same line", func(t *testing.T) {
		assert.NotEqual(t, uint8(0), getPart([]int{1, 2}, 1, []string{
			"...",
			"@0.",
			"...",
		}).symbol)
		assert.NotEqual(t, uint8(0), getPart([]int{1, 2}, 1, []string{
			"...",
			".0@",
			"...",
		}).symbol)
	})

	t.Run("missing", func(t *testing.T) {
		assert.Equal(t, uint8(0), getPart([]int{1, 2}, 1, []string{
			"...",
			".0.",
			"...",
		}).symbol)
	})
}

func TestSolve(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....s
......755.
...$.*....
.664.598..`
	assert.Equal(t, 4361, lo.Must(solve(strings.NewReader(input))))
}
