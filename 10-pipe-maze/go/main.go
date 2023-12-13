package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

func solveFile(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return solve2(f)
}

func solve(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		lines = append(lines, line)
	}
	length := loopLength(lines)
	return length / 2, nil
}

var pipes = map[byte][]Step{
	'S': {{0, 1}, {0, -1}, {1, 0}, {-1, 0}},
	'|': {{0, -1}, {0, 1}},
	'-': {{-1, 0}, {1, 0}},
	'L': {{0, -1}, {1, 0}},
	'J': {{0, -1}, {-1, 0}},
	'7': {{0, 1}, {-1, 0}},
	'F': {{0, 1}, {1, 0}},
	'.': {},
}

var unicodePipes = map[byte]rune{
	'S': 'S',
	'|': '\u2503',
	'-': '\u2501',
	'L': '\u2517',
	'J': '\u251B',
	'7': '\u2513',
	'F': '\u250F',
}

func solve2(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		lines = append(lines, line)
	}
	count := findInsideArea(lines)
	return count, nil
}

type Point struct {
	X, Y int
}

type Step struct {
	X, Y int
}

func (p Point) Step(step Step) Point {
	return Point{p.X + step.X, p.Y + step.Y}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

type Polygon []Point

func (p Polygon) Vertices() Polygon {
	vertices := make(Polygon, 0)
	for i := range p {
		left := p[(i+len(p)-1)%len(p)]
		right := p[(i+1)%len(p)]
		if left.X == right.X || left.Y == right.Y {
			continue
		}
		vertices = append(vertices, p[i])
	}
	return vertices
}

func (p Polygon) Area() int {
	sum := 0
	for i := range p {
		j := (i + 1) % len(p)
		sum += p[i].X * p[j].Y
		sum -= p[j].X * p[i].Y
	}
	sum -= len(p)
	return sum / 2
}

func findInsideArea(lines []string) int {
	poly := getLoop(lines)
	points := poly.Vertices()
	slices.Reverse(points) // TODO: needed?
	fmt.Println(points)
	return points.Area()
}

func printBoard(lines []string, loop map[Point]bool, insides map[Point]bool) {
	for y := range lines {
		for x := range lines[y] {
			if loop[Point{x, y}] {
				fmt.Printf("%c", unicodePipes[lines[y][x]])
				// fmt.Printf("%c", '\u2588')
			} else if insides[Point{x, y}] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func getLoop(lines []string) Polygon {
	start := findStart(lines)
	current := start
	visited := make(map[Point]bool)
	points := make(Polygon, 0)
	length := 0
	for {
		visited[current] = true
		points = append(points, current)
		foundStart := false
		found := false
		for _, step := range pipes[lines[current.Y][current.X]] {
			next := current.Step(step)
			if start == next {
				foundStart = true
				continue
			}
			if visited[next] {
				continue
			}
			// fmt.Printf("current: %c (%d, %d) len(%d) -- try: (%d, %d)\n", lines[y][x], x, y, length, next_x, next_y)
			if canReach(lines, next, current) {
				current = next
				length += 1
				found = true
				// fmt.Printf("found: %c (%d, %d) len(%d)\n", lines[y][x], x, y, length)
				break
			}
		}
		if found {
			continue
		}
		if foundStart {
			return points
		}
	}
}

func loopLength(lines []string) int {
	return len(getLoop(lines))
}

func canReach(lines []string, current, previous Point) bool {
	if current.X < 0 || current.Y < 0 || current.Y >= len(lines) || current.X >= len(lines[current.Y]) {
		return false
	}
	steps := pipes[lines[current.Y][current.X]]
	for _, step := range steps {
		if current.Step(step) == previous {
			return true
		}
	}
	return false
}

func findStart(lines []string) Point {
	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == 'S' {
				return Point{x, y}
			}
		}
	}
	panic("no start found")
}

func main() {
	val, err := solveFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(val)
}
