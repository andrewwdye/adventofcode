package pkg

import (
	"bufio"
	"io"
)

func Solve1(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	lines = tiltUp(lines)
	return weight(lines), nil
}

func Solve2(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	lines = spin(lines, 1000000000)
	return weight(lines), nil
}

func spin(lines []string, times int) []string {
	return lines
}

func weight(lines []string) int {
	total := 0
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == 'O' {
				total += len(lines) - i
			}
		}
	}
	return total
}

func tiltUp(lines []string) []string {
	for j := 0; j < len(lines[0]); j += 1 {
		for i, next := 0, 0; i < len(lines); i += 1 {
			switch lines[i][j] {
			case '.':
				_ = 0
			case 'O':
				if next != i {
					lines[next] = lines[next][:j] + "O" + lines[next][j+1:]
					lines[i] = lines[i][:j] + "." + lines[i][j+1:]
				}
				next += 1
			case '#':
				next = i + 1
			}
		}

	}
	return lines
}
