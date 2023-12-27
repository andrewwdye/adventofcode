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
