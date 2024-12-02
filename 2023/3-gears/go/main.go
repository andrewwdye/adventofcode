package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"

	"github.com/samber/lo"
)

func solveFile(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	return solve(f)
}

func solve(r io.ReadSeeker) (int, error) {
	lines := getLines(r)
	parts := findParts(lines)
	sum := 0
	for _, part := range parts {
		value, err := strconv.Atoi(part.id)
		if err != nil {
			return 0, err
		}
		sum += value
	}
	fmt.Printf("%d\n", sum)

	gear_sum := 0
	for _, gear := range getGears(parts, lines) {
		gear_sum += lo.Must(strconv.Atoi(gear[0])) * lo.Must(strconv.Atoi(gear[1]))
	}
	fmt.Printf("%d\n", gear_sum)
	return sum, nil
}

func getGears(parts []part, lines []string) [][]string {
	gears := make(map[symbol][]string)
	for _, part := range parts {
		sym := lines[part.symbol.line][part.symbol.offset]
		if sym == '*' {
			if _, ok := gears[part.symbol]; !ok {
				gears[part.symbol] = []string{}
			}
			gears[part.symbol] = append(gears[part.symbol], part.id)
		}
	}

	result := [][]string{}
	for _, gear := range gears {
		if len(gear) == 2 {
			result = append(result, gear)
		}
	}
	return result
}

func getLines(r io.ReadSeeker) []string {
	r.Seek(0, 0)
	lines := []string{}
	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

type part struct {
	id     string
	symbol symbol
}

func (p part) valid() bool {
	return p.symbol.line != -1
}

type symbol struct {
	line   int
	offset int
}

func findParts(lines []string) []part {
	parts := []part{}
	re := regexp.MustCompile(`\d+`)
	for i, line := range lines {
		matches := re.FindAllStringIndex(line, -1)
		for _, m := range matches {
			part := getPart(m, i, lines)
			if part.valid() {
				parts = append(parts, part)
			}
		}
	}
	return parts
}

func getPart(m []int, line int, lines []string) part {
	id := lines[line][m[0]:m[1]]
	// Check line above
	for i := m[0] - 1; i <= m[1]; i++ {
		if isLocSymbol(line-1, i, lines) {
			return part{
				id:     id,
				symbol: symbol{line - 1, i},
			}
		}
	}

	// Check line
	if isLocSymbol(line, m[0]-1, lines) {
		return part{
			id:     id,
			symbol: symbol{line, m[0] - 1},
		}
	}
	if isLocSymbol(line, m[1], lines) {
		return part{
			id:     id,
			symbol: symbol{line, m[1]},
		}
	}

	// Check line below
	for i := m[0] - 1; i <= m[1]; i++ {
		if isLocSymbol(line+1, i, lines) {
			return part{
				id:     id,
				symbol: symbol{line + 1, i},
			}
		}
	}
	return part{
		id:     id,
		symbol: symbol{-1, -1},
	}
}

func isLocSymbol(line, offset int, lines []string) bool {
	if line < 0 || line >= len(lines) || offset < 0 || offset >= len(lines[line]) {
		return false
	}
	c := lines[line][offset]
	if c == '.' || (c >= '0' && c <= '9') {
		return false
	}
	return true
}

func main() {
	_, err := solveFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
