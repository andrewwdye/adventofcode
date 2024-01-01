package pkg

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Direction string

const (
	Up    Direction = "U"
	Down  Direction = "D"
	Left  Direction = "L"
	Right Direction = "R"
)

func (d Direction) Opposite() Direction {
	switch d {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	}
	panic("invalid direction")
}

func (d Direction) String() string {
	switch d {
	case Up:
		return "^"
	case Down:
		return "v"
	case Left:
		return "<"
	case Right:
		return ">"
	}
	panic("invalid direction")
}

type Vertex struct {
	X, Y     int
	Dir      Direction
	Distance int
}

func (v Vertex) Move(dir Direction, distance int) Vertex {
	switch dir {
	case Up:
		return Vertex{v.X, v.Y - distance, dir, distance}
	case Down:
		return Vertex{v.X, v.Y + distance, dir, distance}
	case Left:
		return Vertex{v.X - distance, v.Y, dir, distance}
	case Right:
		return Vertex{v.X + distance, v.Y, dir, distance}
	}
	panic("invalid direction")
}

type Polygon struct {
	Vertices         []Vertex
	Perimeter        int
	Width, Height    int
	OffsetX, OffsetY int
}

func (p *Polygon) Area() int {
	area := 0
	for i := 0; i < len(p.Vertices)-1; i++ {
		j := (i + 1) % len(p.Vertices)
		area += p.Vertices[i].X*p.Vertices[j].Y - p.Vertices[j].X*p.Vertices[i].Y
	}
	if area < 0 {
		area = -area
	}
	return area/2 + p.Perimeter/2 + 1
	//  ## 0,0 1,0 p=4
	//  ## 1,0 1,1
	//
}

func (p *Polygon) String() string {
	out := make([]string, 0, p.Height)
	for y := 0; y < p.Height; y++ {
		out = append(out, strings.Repeat(".", p.Width))
	}
	for _, v := range p.Vertices {
		v = Vertex{v.X + p.OffsetX, v.Y + p.OffsetY, v.Dir, v.Distance}
		s := []byte(out[v.Y])
		s[v.X] = v.Dir.Opposite().String()[0]
		out[v.Y] = string(s)

	}
	return strings.Join(out, "\n")
}

func NewPolygon(reader io.Reader) Polygon {
	scanner := bufio.NewScanner(reader)
	re := regexp.MustCompile(`([U|D|L|R]) ([0-9]+) \(#([0-9a-f]{6})\)`)
	polygon := Polygon{}
	curr := Vertex{0, 0, Up, 0}
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) != 4 {
			panic(fmt.Sprint("invalid line", line))
		}
		count := lo.Must(strconv.Atoi(matches[2]))
		polygon.Perimeter += count
		curr = curr.Move(Direction(matches[1]), count)
		if curr.X < minX {
			minX = curr.X
		}
		if curr.Y < minY {
			minY = curr.Y
		}
		if curr.X > maxX {
			maxX = curr.X
		}
		if curr.Y > maxY {
			maxY = curr.Y
		}
		polygon.Vertices = append(polygon.Vertices, curr)
	}
	polygon.Width = maxX - minX + 1
	polygon.Height = maxY - minY + 1
	polygon.OffsetX = -minX
	polygon.OffsetY = -minY
	return polygon
}

func Solve1(reader io.Reader) (int, error) {
	poly := NewPolygon(reader)
	// fmt.Println(poly)
	// fmt.Println(poly.String())
	return poly.Area(), nil
}
