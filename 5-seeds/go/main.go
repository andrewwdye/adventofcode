package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/samber/lo"
)

func parseFile(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	total, err := parse(f)
	if err != nil {
		return 0, err
	}
	fmt.Println(total)
	return total, nil
}

func parse(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)

	spans := getSeedSpans(scanner)
	maps := []ResourceMap{
		NewResourceMap(scanner), // to soil
		NewResourceMap(scanner), // to fertilizer
		NewResourceMap(scanner), // to water
		NewResourceMap(scanner), // to light
		NewResourceMap(scanner), // to temp
		NewResourceMap(scanner), // to humidity
		NewResourceMap(scanner), // to location
	}

	min := math.MaxInt32
	for i, span := range spans {
		fmt.Printf("span: %d\n", i)
		for seed := span.start; seed < span.start+span.span; seed++ {
			v := seed
			for _, m := range maps {
				v = m.Get(v)
			}
			if v < min {
				min = v
			}
		}
	}
	return min, nil
}

// func getSeeds(scanner *bufio.Scanner) []int {
// 	re := regexp.MustCompile(`\d+`)
// 	scanner.Scan()
// 	seeds := lo.Map(re.FindAllString(scanner.Text(), -1), func(item string, _ int) int {
// 		return lo.Must(strconv.Atoi(item))
// 	})
// 	scanner.Scan()
// 	scanner.Text()
// 	return seeds
// }

type SeedSpan struct {
	start int
	span  int
}

func getSeedSpans(scanner *bufio.Scanner) []SeedSpan {
	re := regexp.MustCompile(`\d+`)
	scanner.Scan()
	seeds := lo.Map(re.FindAllString(scanner.Text(), -1), func(item string, _ int) int {
		return lo.Must(strconv.Atoi(item))
	})
	scanner.Scan()
	scanner.Text()

	seedSpans := make([]SeedSpan, 0)
	i := 0
	for i < len(seeds) {
		seedSpans = append(seedSpans, SeedSpan{seeds[i], seeds[i+1]})
		i += 2
	}

	return seedSpans
}

type ResourceMapEntry struct {
	source      int
	destination int
	span        int
}

type ResourceMap []ResourceMapEntry

func (r ResourceMap) Get(source int) int {
	for _, entry := range r {
		if source < entry.source {
			break
		} else if source >= entry.source && source < entry.source+entry.span {
			return entry.destination + source - entry.source
		}
	}
	return source
}

func NewResourceMap(scanner *bufio.Scanner) ResourceMap {
	result := make(ResourceMap, 0)
	scanner.Scan()
	scanner.Text()
	re := regexp.MustCompile(`\d+`)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		values := lo.Map(re.FindAllString(line, -1), func(item string, _ int) int {
			return lo.Must(strconv.Atoi(item))
		})
		if len(values) != 3 {
			panic(fmt.Sprintf("invalid line %s", line))
		}
		result = append(result, ResourceMapEntry{
			source:      values[1],
			destination: values[0],
			span:        values[2],
		})
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].source < result[j].source
	})
	return result
}

func main() {
	_, err := parseFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
