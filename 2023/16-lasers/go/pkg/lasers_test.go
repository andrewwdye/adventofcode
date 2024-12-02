package pkg

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestSolve1(t *testing.T) {
	input := `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

	assert.Equal(t, 46, lo.Must(Solve1(strings.NewReader(input))))
}

func TestSolve2(t *testing.T) {
	input := `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

	assert.Equal(t, 51, lo.Must(Solve2(strings.NewReader(input))))
}

func TestRun(t *testing.T) {
	input := `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

	state1 := NewState(strings.NewReader(input))
	assert.Equal(t, 46, state1.Run(Laser{0, 0, Right}))

	state2 := NewState(strings.NewReader(input))
	assert.Equal(t, 51, state2.Run(Laser{3, 0, Down}))
}
