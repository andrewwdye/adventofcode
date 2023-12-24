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
	Current     byte
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
	// Append a '.' so we can assume the record end in a non-spring
	return countWaysRecurse(record+".", 0, ways, 0, cache), nil
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
	if count, ok := cache[PrecomputedEntry{offset, waysOffset, c}]; ok {
		// fmt.Printf("cache hit offset: %d, waysOffset: %d, ways:%v, count: %d\n", offset, waysOffset, ways[waysOffset:], count)
		return count
	}
	count := 0
	switch c {
	case '.':
		count = countWaysRecurse(record, offset+1, ways, waysOffset, cache)
	case '#':
		if findStreak(record, offset, ways, waysOffset) {
			count = countWaysRecurse(record, offset+ways[waysOffset]+1, ways, waysOffset+1, cache)
		} else {
			count = 0
		}
	case '?':
		count = countWaysRecurse(record[:offset]+"#"+record[offset+1:], offset, ways, waysOffset, cache) +
			countWaysRecurse(record[:offset]+"."+record[offset+1:], offset, ways, waysOffset, cache)
	default:
		panic(fmt.Sprintf("invalid char: %c", c))
	}
	cache[PrecomputedEntry{offset, waysOffset, c}] = count
	return count
}

func findStreak(record string, offset int, ways []int, waysOffset int) bool {
	if waysOffset >= len(ways) {
		return false
	}
	streak := ways[waysOffset]
	if streak > len(record)-offset {
		return false
	}
	for _, c := range record[offset : offset+streak] {
		if c == '.' {
			return false
		}
	}
	// Record will always end in an '.'
	if record[offset+streak] == '#' {
		return false
	}
	return true
}
