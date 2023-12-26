package pkg

import (
	"bufio"
	"io"
)

func Solve(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return tilt(lines), nil
}

func tilt(lines []string) int {
	total := 0
	for i := 0; i < len(lines[0]); i += 1 {
		for j, next := 0, 0; j < len(lines); j += 1 {
			switch lines[j][i] {
			case '.':
				_ = 0
			case 'O':
				total += len(lines) - next
				next += 1
			case '#':
				next = j + 1
			}
		}

	}
	return total
}
