package pkg

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridSearch(t *testing.T) {
	input := `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

	g := NewGrid(strings.NewReader(input), 1, 3)
	assert.Equal(t, 102, g.Search())
}

func TestGridSearch2(t *testing.T) {
	t.Run("original", func(t *testing.T) {
		input := `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

		g := NewGrid(strings.NewReader(input), 4, 10)
		assert.Equal(t, 94, g.Search())
	})
	t.Run("small", func(t *testing.T) {
		input := `111111111111
999999999991
999999999991
999999999991
999999999991`

		g := NewGrid(strings.NewReader(input), 4, 10)
		assert.Equal(t, 71, g.Search())
	})
}
