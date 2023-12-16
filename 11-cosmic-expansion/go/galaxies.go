package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/samber/lo"
)

func solve1(reader io.Reader) (int, error) {
	_ = ParseUniverse(reader)
	return 0, nil
}

type Location struct {
	X, Y int
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

	// insert rows
	expandedLines := []string{}
	for _, line := range lines {
		expandedLines = append(expandedLines, line)
		if lo.EveryBy([]byte(line), func(b byte) bool {
			return b == '.'
		}) {
			expandedLines = append(expandedLines, line)
		}
	}

	// insert columns
	for i := range lines[0] {
		if lo.EveryBy(lines, func(line string) bool {
			return line[i] == '.'
		}) {
			for j := range expandedLines {
				expandedLines[j] = expandedLines[j][:i] + expandedLines[j][i:i+1] + expandedLines[j][i:]
			}
		}
	}

	universe := Universe{
		Width:    len(expandedLines[0]),
		Height:   len(expandedLines),
		Original: lines,
		Expanded: expandedLines,
	}
	for row, line := range expandedLines {
		for col := range line {
			if line[col] == '#' {
				universe.Galaxies = append(universe.Galaxies, Location{col, row})
			}
		}
	}
	return universe
}
