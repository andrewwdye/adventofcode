package pkg

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestSolve1(t *testing.T) {
	input := `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`

	assert.Equal(t, 62, lo.Must(Solve1(strings.NewReader(input))))
}

func TestArea(t *testing.T) {
	p1 := Polygon{Vertices: []Vertex{{0, 0, Up, 0}, {0, 1, Up, 0}, {1, 1, Up, 0}, {1, 0, Up, 0}}, Perimeter: 4}
	assert.Equal(t, 4, p1.Area())

	p2 := Polygon{Vertices: []Vertex{{0, 0, Up, 0}, {0, 2, Up, 0}, {2, 2, Up, 0}, {2, 0, Up, 0}}, Perimeter: 8}
	assert.Equal(t, 9, p2.Area())
}
