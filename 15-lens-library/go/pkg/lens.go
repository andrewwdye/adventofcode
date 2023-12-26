package pkg

import (
	"bufio"
	"io"
	"strings"
)

func Solve1(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		for _, seq := range strings.Split(line, ",") {
			sum += hash(seq)
		}
	}
	return sum, nil
}

func hash(input string) int {
	current := 0
	for _, c := range input {
		current = ((current + int(c)) * 17) % 256
	}
	return current
}
