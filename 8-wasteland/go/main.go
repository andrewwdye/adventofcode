package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/samber/lo"
)

func solveFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return solve(file)
}

func solve(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	instructions := scanner.Text()

	scanner.Scan()

	re := regexp.MustCompile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`)
	network := map[string][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) != 4 {
			return 0, fmt.Errorf("invalid input: %s", line)
		}
		network[matches[1]] = []string{matches[2], matches[3]}
	}
	return walk(instructions, network)
}

func walk(instructions string, network map[string][]string) (int, error) {
	count := 0
	current := "AAA"
	for {
		if current == "ZZZ" {
			return count, nil
		}
		dir := lo.Ternary(instructions[count%len(instructions)] == 'L', 0, 1)
		current = network[current][dir]
		count += 1
	}
}

func main() {
	count, err := solveFile(os.Args[1])
	fmt.Println(count)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
