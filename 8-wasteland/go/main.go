package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"

	"github.com/samber/lo"
)

func solveFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return solve(file, true)
}

func solve(reader io.Reader, ghosts bool) (int, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	instructions := scanner.Text()

	scanner.Scan()

	re := regexp.MustCompile(`([A-Z0-9]{3}) = \(([A-Z0-9]{3}), ([A-Z0-9]{3})\)`)
	network := map[string][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) != 4 {
			return 0, fmt.Errorf("invalid input: %s", line)
		}
		network[matches[1]] = []string{matches[2], matches[3]}
	}
	if ghosts {
		return walkGhosts(instructions, network)
	} else {
		return walk(instructions, network)
	}
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

func walkGhosts(instructions string, network map[string][]string) (int, error) {
	starts := []string{}
	for name := range network {
		if name[2] == 'A' {
			starts = append(starts, name)
		}
	}
	counts := make([]int, len(starts))
	for i, start := range starts {
		count := 0
		current := start
		for {
			if current[2] == 'Z' {
				counts[i] = count
				fmt.Printf("%d: %d\n", i, count)
				break
			}
			dir := lo.Ternary(instructions[count%len(instructions)] == 'L', 0, 1)
			current = network[current][dir]
			count += 1
		}
	}
	fmt.Println(counts)
	result := lcm(counts[0], counts[1], counts[2:]...)
	// for _, count := range counts {
	// 	fmt.Println(result % count)
	// }
	return result, nil
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)
	for _, integer := range integers {
		result = lcm(result, integer)
	}
	return result
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	nums := []int{b, a % b}
	slices.Sort(nums)
	return gcd(nums[1], nums[0])
}

func main() {
	count, err := solveFile(os.Args[1])
	fmt.Println(count)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
