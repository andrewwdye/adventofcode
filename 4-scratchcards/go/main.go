package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
)

func parseFile(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	total, err := parse2(f)
	if err != nil {
		return 0, err
	}
	fmt.Println(total)
	return total, nil
}

func parse(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		value, err := parseLine(line)
		if err != nil {
			return 0, err
		}
		total += value
	}
	return total, nil
}

func parseLine(line string) (int, error) {
	line_re := regexp.MustCompile(`.+:(.+)\|(.+)`)
	matches := line_re.FindStringSubmatch(line)
	if len(matches) != 3 {
		return 0, fmt.Errorf("invalid line: %s", line)
	}
	num_re := regexp.MustCompile(`\d+`)
	winners := make(map[string]bool)

	for _, num := range num_re.FindAllString(matches[1], -1) {
		winners[num] = true
	}
	found := 0
	for _, num := range num_re.FindAllString(matches[2], -1) {
		if _, ok := winners[num]; ok {
			found++
		}
	}
	if found == 0 {
		return 0, nil
	}
	return int(math.Pow(2, float64(found-1))), nil
}

func parse2(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	cards := make(map[int]int)
	card := 0
	for scanner.Scan() {
		cards[card] += 1
		line := scanner.Text()
		value, err := parseLine2(line)
		if err != nil {
			return 0, err
		}
		for i := 0; i < value; i++ {
			cards[card+i+1] += cards[card]
		}
		card += 1
	}
	total := 0
	for _, v := range cards {
		total += v
	}
	return total, nil
}

func parseLine2(line string) (int, error) {
	line_re := regexp.MustCompile(`.+:(.+)\|(.+)`)
	matches := line_re.FindStringSubmatch(line)
	if len(matches) != 3 {
		return 0, fmt.Errorf("invalid line: %s", line)
	}
	num_re := regexp.MustCompile(`\d+`)
	winners := make(map[string]bool)

	for _, num := range num_re.FindAllString(matches[1], -1) {
		winners[num] = true
	}
	found := 0
	for _, num := range num_re.FindAllString(matches[2], -1) {
		if _, ok := winners[num]; ok {
			found++
		}
	}
	return found, nil
}

func main() {
	_, err := parseFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
