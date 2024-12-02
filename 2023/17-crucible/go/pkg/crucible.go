package pkg

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"

	"github.com/samber/lo"
)

type Location struct {
	X, Y int
}

func (l Location) Step(s Step) Location {
	return Location{l.X + s.X, l.Y + s.Y}
}

func (l Location) ManhattanDist(other Location) int {
	dx := l.X - other.X
	dy := l.Y - other.Y
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

type Step struct {
	X, Y int
}

type Direction int

func (d Direction) Step() Step {
	switch d {
	case Up:
		return Step{0, -1}
	case Down:
		return Step{0, 1}
	case Left:
		return Step{-1, 0}
	case Right:
		return Step{1, 0}
	}
	panic("invalid direction")
}

func (d Direction) Opposite() Direction {
	switch d {
	case Unknown:
		return Unknown
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
	case Unknown:
		return " "
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

const (
	Unknown Direction = iota
	Up
	Down
	Left
	Right
)

type Node struct {
	Location
	ArrivalDir Direction
	Streak     int
}

type NodeLoss struct {
	Node
	MinLossToHere int
}

type Grid struct {
	Losses    [][]int
	Mins      map[Location]int
	Dirs      map[Location]Direction
	MinStreak int
	MaxStreak int
}

func (g Grid) SprintLosses() string {
	out := ""
	for y := range g.Losses {
		for x := range g.Losses[y] {
			out += fmt.Sprintf("%d", g.Losses[y][x])
		}
		out += "\n"
	}
	return out
}

func (g Grid) SprintMins() string {
	out := ""
	for y := range g.Losses {
		for x := range g.Losses[y] {
			out += fmt.Sprintf("%4d", g.Mins[Location{x, y}])
		}
		out += "\n"
	}
	return out
}

func (g Grid) SprintDirs() string {
	out := ""
	for y := range g.Losses {
		for x := range g.Losses[y] {
			out += g.Dirs[Location{x, y}].String()
		}
		out += "\n"
	}
	return out
}

func (g Grid) Start() Location {
	return Location{0, 0}
}

func (g Grid) End() Location {
	return Location{len(g.Losses[0]) - 1, len(g.Losses) - 1}
}

func (g Grid) Search() int {
	// Keep track of the nodes we haven't found yet
	available := map[Node]bool{}
	for y := range g.Losses {
		for x := range g.Losses[y] {
			for _, dir := range []Direction{Up, Down, Left, Right} {
				for streak := 1; streak <= g.MaxStreak; streak += 1 {
					available[Node{Location{x, y}, dir, streak}] = true
				}
			}
		}
	}
	// Nodes we've found but haven't committed yet
	working := &NodeHeap{Target: g.End()}
	heap.Init(working)
	heap.Push(working, NodeLoss{Node{g.Start(), Unknown, 0}, 0})

	for working.Len() > 0 {
		node := heap.Pop(working).(NodeLoss)
		if _, ok := g.Mins[node.Location]; !ok || node.MinLossToHere < g.Mins[node.Location] {
			g.Mins[node.Location] = node.MinLossToHere
			g.Dirs[node.Location] = node.ArrivalDir
		}
		if node.Location == g.End() && node.Streak >= g.MinStreak {
			// fmt.Println(g.SprintLosses())
			// fmt.Println(g.SprintMins())
			// fmt.Println(g.SprintDirs())
			return node.MinLossToHere
		}
		// fmt.Printf("from (%d,%d) dir: %v, streak: %d, loss: %d\n", node.X, node.Y, node.ArrivalDir, node.Streak, node.MinLossToHere)
		for _, dir := range []Direction{Up, Down, Left, Right} {
			if dir == node.ArrivalDir.Opposite() {
				continue
			}
			if dir == node.ArrivalDir && node.Streak > g.MaxStreak {
				continue
			}
			if dir != node.ArrivalDir && (node.Streak > 0 && node.Streak < g.MinStreak) {
				continue
			}
			nextLoc := node.Location.Step(dir.Step())
			next := Node{nextLoc, dir, 1}
			if dir == node.ArrivalDir {
				next.Streak = node.Streak + 1
			}
			if !available[next] {
				continue
			}
			available[next] = false
			nextLoss := NodeLoss{next, node.MinLossToHere + g.Losses[nextLoc.Y][nextLoc.X]}
			// fmt.Printf("  to (%d,%d) dir: %v, streak: %d, loss: %d\n", nextLoss.X, nextLoss.Y, nextLoss.ArrivalDir, nextLoss.Streak, nextLoss.MinLossToHere)
			heap.Push(working, nextLoss)
		}
	}
	panic("no path found")
}

func NewGrid(reader io.Reader, min, max int) Grid {
	scanner := bufio.NewScanner(reader)
	grid := Grid{MinStreak: min, MaxStreak: max}
	for scanner.Scan() {
		line := scanner.Text()
		row := lo.Map([]byte(line), func(c byte, _ int) int {
			return int(c) - '0'
		})
		grid.Losses = append(grid.Losses, row)
		grid.Mins = map[Location]int{}
		grid.Dirs = map[Location]Direction{}
	}
	return grid
}

func Solve1(reader io.Reader) (int, error) {
	return NewGrid(reader, 1, 3).Search(), nil
}

func Solve2(reader io.Reader) (int, error) {
	return NewGrid(reader, 4, 10).Search(), nil
}
