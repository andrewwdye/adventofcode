package main

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	input := `.....
.S-7.
.|.|.
.L-J.
.....`
	assert.Equal(t, 4, lo.Must(solve(strings.NewReader(input))))
}

func TestSolve2(t *testing.T) {
	t.Run("sample1", func(t *testing.T) {
		input := `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`
		assert.Equal(t, 4, lo.Must(solve2(strings.NewReader(input))))
	})

	t.Run("sample2", func(t *testing.T) {
		input := `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`
		assert.Equal(t, 8, lo.Must(solve2(strings.NewReader(input))))
	})

	t.Run("sample3", func(t *testing.T) {
		input := `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`
		assert.Equal(t, 10, lo.Must(solve2(strings.NewReader(input))))
	})
}
