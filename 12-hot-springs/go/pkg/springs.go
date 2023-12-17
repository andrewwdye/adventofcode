package pkg

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type PrecomputedEntry struct {
	RecordIndex int
	WaysIndex   int
}

func Solve(reader io.Reader, expand bool) (int, error) {
	scanner := bufio.NewScanner(reader)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		ways, err := countWays(line, expand)
		if err != nil {
			return 0, err
		}
		sum += ways
	}
	return sum, nil
}

func countWays(line string, expand bool) (int, error) {
	re := regexp.MustCompile(`([.#?]+) ([\d,]+)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) != 3 {
		panic("invalid line")
	}
	record := matches[1]
	ways := lo.Map(strings.Split(matches[2], ","), func(s string, _ int) int {
		return lo.Must(strconv.Atoi(s))
	})

	if expand {
		expandedRecord := record
		for i := 0; i < 4; i++ {
			expandedRecord += "?" + record
		}
		record = expandedRecord

		expandedWays := make([]int, 0, len(ways)*5)
		for i := 0; i < 5; i++ {
			expandedWays = append(expandedWays, ways...)
		}
		ways = expandedWays
	}

	fmt.Println(record, ways)
	cache := make(map[PrecomputedEntry]int)
	return countWaysRecurse(record, 0, ways, 0, cache), nil
}

func countWaysRecurse(record string, offset int, ways []int, waysOffset int, cache map[PrecomputedEntry]int) int {
	if offset >= len(record) {
		if waysOffset >= len(ways) {
			// fmt.Printf("success: %s\n", record)
			return 1
		} else {
			return 0
		}
	}
	c := record[offset]
	switch c {
	case '.':
		return countWaysRecurse(record, offset+1, ways, waysOffset, cache)
	case '#':
		if count, ok := findStreak(record, offset, ways, waysOffset); ok {
			return countWaysRecurse(record, offset+count, ways, waysOffset+1, cache)
		} else {
			return 0
		}
	case '?':
		return countWaysRecurse(record[:offset]+"#"+record[offset+1:], offset, ways, waysOffset, cache) +
			countWaysRecurse(record[:offset]+"."+record[offset+1:], offset, ways, waysOffset, cache)
	default:
		panic(fmt.Sprintf("invalid char: %c", c))
	}
}

func findStreak(record string, offset int, ways []int, waysOffset int) (int, bool) {
	if waysOffset >= len(ways) {
		return 0, false
	}
	streak := ways[waysOffset]
	if streak > len(record)-offset {
		return 0, false
	}
	for _, c := range record[offset : offset+streak] {
		if c == '.' {
			return 0, false
		}
	}
	if len(record)-offset == streak {
		return streak, true
	}
	if record[offset+streak] == '#' {
		return 0, false
	}
	return streak + 1, true
}

// func countWaysRecurse(record string, offset int, ways []int, waysOffset int, streak int, cache map[PrecomputedEntry]int) int {
// 	// TOOD: check valid streaks
// 	if offset >= len(record) && waysOffset < len(ways) {
// 		return 0
// 	}
// 	//...
// 	// Check cache
// 	if waysOffset < len(ways) && streak == 0 {
// 		if count, ok := cache[PrecomputedEntry{offset, ways[waysOffset]}]; ok {
// 			fmt.Printf("success: %s\n", record)
// 			return count
// 		}
// 	}
// 	// We are done if remaining springs <= expected
// 	known := 0
// 	for i := offset; i < len(record); i++ {
// 		if record[i] == '#' {
// 			known += 1
// 		}
// 	}
// 	expected := -offset
// 	for _, way := range ways[waysOffset:] {
// 		expected += way
// 	}
// 	if known == expected {
// 		fmt.Printf("success: %s\n", record)
// 		return 1
// 	} else if known > expected {
// 		return 0
// 	}
// 	// Recurse
// 	c := record[offset]
// 	switch c {
// 	case '.':
// 		if streak > 0 {
// 			if ways[waysOffset] != streak {
// 				return 0
// 			}
// 			return countWaysRecurse(record, offset+1, ways, waysOffset+1, 0, cache)
// 		}
// 		return countWaysRecurse(record, offset+1, ways, waysOffset, 0, cache)
// 	case '#':
// 		if waysOffset >= len(ways) {
// 			return 0
// 		}
// 		streak += 1
// 		if streak > ways[waysOffset] {
// 			return 0
// 		}
// 		return countWaysRecurse(record, offset+1, ways, waysOffset, streak, cache)
// 	case '?':
// 		return countWaysRecurse(record[:offset]+"#"+record[offset+1:], offset, ways, waysOffset, streak, cache) +
// 			countWaysRecurse(record[:offset]+"."+record[offset+1:], offset, ways, waysOffset, streak, cache)
// 	default:
// 		panic(fmt.Sprintf("invalid char: %c", c))
// 	}
// }
