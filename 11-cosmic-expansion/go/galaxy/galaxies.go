package galaxy

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/samber/lo"
)

func Solve1(reader io.Reader) (int, error) {
	univ := ParseUniverse(reader)
	sum := 0
	for i, galaxy := range univ.Galaxies {
		for _, other := range univ.Galaxies[i+1:] {
			sum += galaxy.Distance(other)
		}
	}
	return sum, nil
}

type Location struct {
	X, Y int
}

func (l Location) Distance(other Location) int {
	dx := l.X - other.X
	if dx < 0 {
		dx = -dx
	}
	dy := l.Y - other.Y
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

func (l Location) String() string {
	return fmt.Sprintf("(%d, %d)", l.X, l.Y)
}

type Universe struct {
	Width, Height int
	Galaxies      []Location
	Original      []string
	Expanded      []string
}

func (u Universe) String() string {
	return strings.Join(u.Expanded, "\n")
}

func ParseUniverse(reader io.Reader) Universe {
	scanner := bufio.NewScanner(reader)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	expanded := expand(lines)
	universe := Universe{
		Width:    len(expanded[0]),
		Height:   len(expanded),
		Original: lines,
		Expanded: expanded,
	}
	for row, line := range expanded {
		for col := range line {
			if line[col] == '#' {
				universe.Galaxies = append(universe.Galaxies, Location{col, row})
			}
		}
	}
	return universe
}

func expand(lines []string) []string {
	expandedLines := []string{}
	for _, line := range lines {
		expandedLines = append(expandedLines, line)
		if lo.EveryBy([]byte(line), func(b byte) bool {
			return b == '.'
		}) {
			expandedLines = append(expandedLines, line)
		}
	}
	expanded := 0
	for i := range lines[0] {
		if lo.EveryBy(lines, func(line string) bool {
			return line[i] == '.'
		}) {
			for j := range expandedLines {
				expandedLines[j] = expandedLines[j][:i+expanded] + "." + expandedLines[j][i+expanded:]
			}
			expanded += 1
		}
	}
	return expandedLines
}
