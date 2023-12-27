package pkg

import (
	"bufio"
	"fmt"
	"io"

	"github.com/samber/lo"
)

const (
	MAX_GRID_ROUNDS   = -1
	PRINT_GRID_PERIOD = -1
)

type Element byte

func (e Element) String() string {
	return string(e)
}

const (
	Empty              Element = '.'
	LeftMirror         Element = '\\'
	RightMirror        Element = '/'
	VerticalSplitter   Element = '|'
	HorizontalSplitter Element = '-'
)

type Cell struct {
	Element   Element
	Energized bool
	entered   map[Dir]bool
}

func (c Cell) Enter(dir Dir) {
	c.entered[dir] = true
}

func (c Cell) EnteredFrom(dir Dir) bool {
	return c.entered[dir]
}

type Grid [][]Cell

func (g Grid) Cell(x, y int) Cell {
	return g[y][x]
}

func (g Grid) Energize(laser Laser) {
	g[laser.y][laser.x].Energized = true
	g[laser.y][laser.x].Enter(laser.dir)
}

func (g Grid) InBounds(x, y int) bool {
	return x >= 0 && x < len(g[0]) && y >= 0 && y < len(g)
}

func (g Grid) ProgressLaser(laser Laser) []Laser {
	g.Energize(laser)
	cell := g.Cell(laser.x, laser.y)
	lasers := []Laser{}
	switch cell.Element {
	case Empty:
		lasers = append(lasers, laser.Next())
	case LeftMirror:
		lasers = append(lasers, laser.LeftMirrorNext())
	case RightMirror:
		lasers = append(lasers, laser.RightMirrorNext())
	case VerticalSplitter:
		lasers = append(lasers, laser.VerticalSplitterNext()...)
	case HorizontalSplitter:
		lasers = append(lasers, laser.HorizontalSplitterNext()...)
	}
	nextLasers := lo.Filter(lasers, func(laser Laser, _ int) bool {
		return g.InBounds(laser.x, laser.y) && !g.Cell(laser.x, laser.y).EnteredFrom(laser.dir)
	})
	// fmt.Printf("In:      %s\n", laser)
	// fmt.Printf("Element: %s\n", element)
	// fmt.Printf("Out:     %v\n", nextLasers)
	return nextLasers
}

func (g Grid) Energy() int {
	energy := 0
	for _, row := range g {
		for _, cell := range row {
			if cell.Energized {
				energy++
			}
		}
	}
	return energy
}

func (g Grid) String() string {
	str := ""
	for _, row := range g {
		for _, cell := range row {
			if cell.Element == Empty {
				if cell.Energized {
					str += "*"
				} else {
					str += "."
				}
			} else {
				str += string(cell.Element)
			}
		}
		str += "\n"
	}
	return str
}

type Dir int

func (d Dir) String() string {
	switch d {
	case Unknown:
		return "Unknown"
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	}
	panic("invalid direction")
}

const (
	Unknown Dir = iota
	Up
	Down
	Left
	Right
)

type Laser struct {
	x, y int
	dir  Dir
}

func (l Laser) String() string {
	return fmt.Sprintf("(%d, %d) %s", l.x, l.y, l.dir)
}

func (l Laser) Next() Laser {
	switch l.dir {
	case Up:
		return Laser{l.x, l.y - 1, l.dir}
	case Down:
		return Laser{l.x, l.y + 1, l.dir}
	case Left:
		return Laser{l.x - 1, l.y, l.dir}
	case Right:
		return Laser{l.x + 1, l.y, l.dir}
	}
	panic("invalid direction")
}

func (l Laser) LeftMirrorNext() Laser {
	switch l.dir {
	case Up:
		return Laser{l.x - 1, l.y, Left}
	case Down:
		return Laser{l.x + 1, l.y, Right}
	case Left:
		return Laser{l.x, l.y - 1, Up}
	case Right:
		return Laser{l.x, l.y + 1, Down}
	}
	panic("invalid direction")
}

func (l Laser) RightMirrorNext() Laser {
	switch l.dir {
	case Up:
		return Laser{l.x + 1, l.y, Right}
	case Down:
		return Laser{l.x - 1, l.y, Left}
	case Left:
		return Laser{l.x, l.y + 1, Down}
	case Right:
		return Laser{l.x, l.y - 1, Up}
	}
	panic("invalid direction")
}

func (l Laser) VerticalSplitterNext() []Laser {
	switch l.dir {
	case Up:
		fallthrough
	case Down:
		return []Laser{l.Next()}
	case Left:
		fallthrough
	case Right:
		return []Laser{l.LeftMirrorNext(), l.RightMirrorNext()}
	}
	panic("invalid direction")
}

func (l Laser) HorizontalSplitterNext() []Laser {
	switch l.dir {
	case Up:
		fallthrough
	case Down:
		return []Laser{l.LeftMirrorNext(), l.RightMirrorNext()}
	case Left:
		fallthrough
	case Right:
		return []Laser{l.Next()}
	}
	panic("invalid direction")
}

type State struct {
	grid   Grid
	lasers []Laser
}

func (s *State) Tick() {
	nextLasers := []Laser{}
	for _, laser := range s.lasers {
		nextLasers = append(nextLasers, s.grid.ProgressLaser(laser)...)
	}
	s.lasers = nextLasers
}

func (s *State) Energy() int {
	return s.grid.Energy()
}

func (s *State) Run(start Laser) int {
	// First laser is always at the top left, moving right
	s.Reset()
	s.lasers = []Laser{start}
	for i := 0; len(s.lasers) > 0 && (MAX_GRID_ROUNDS < 0 || i < MAX_GRID_ROUNDS); i++ {
		if PRINT_GRID_PERIOD > 0 && i%PRINT_GRID_PERIOD == 0 {
			fmt.Println(s.grid)
		}
		s.Tick()
	}
	energy := s.Energy()
	// fmt.Println(start, energy)
	return energy
}

func (s *State) Run2() int {
	max := 0
	for x := 0; x < len(s.grid[0]); x++ {
		energy := s.Run(Laser{x, 0, Down})
		if energy > max {
			max = energy
		}
		energy = s.Run(Laser{x, len(s.grid) - 1, Up})
		if energy > max {
			max = energy
		}
	}
	for y := 0; y < len(s.grid); y++ {
		energy := s.Run(Laser{0, y, Right})
		if energy > max {
			max = energy
		}
		energy = s.Run(Laser{len(s.grid[0]) - 1, y, Left})
		if energy > max {
			max = energy
		}
	}
	return max
}

func (s *State) Reset() {
	for y := range s.grid {
		for x := range s.grid[y] {
			s.grid[y][x].Energized = false
			s.grid[y][x].entered = map[Dir]bool{}
		}
	}
}

func NewState(reader io.Reader) State {
	scanner := bufio.NewScanner(reader)
	state := State{}
	for scanner.Scan() {
		line := scanner.Text()
		cells := lo.Map([]byte(line), func(c byte, _ int) Cell {
			return Cell{Element(c), false, map[Dir]bool{}}
		})
		state.grid = append(state.grid, cells)
	}
	return state
}

func Solve1(reader io.Reader) (int, error) {
	state := NewState(reader)
	return state.Run(Laser{x: 0, y: 0, dir: Right}), nil
}

func Solve2(reader io.Reader) (int, error) {
	state := NewState(reader)
	return state.Run2(), nil
}
