package pkg

import (
	"bufio"
	"io"
	"slices"
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

func encodeLines(lines []string) string {
	result := ""
	for _, line := range lines {
		result += line
	}
	return result
}

func copyLines(lines []string) []string {
	out := make([]string, 0, len(lines))
	return append(out, lines...)
}

func spin(lines []string, times int) []string {
	cache := map[string][]string{}
	firstSeen := map[string]int{}
	// w := make([]int, 100)
	for i := 0; i < times; i += 1 {
		// w = append(w[1:100], weight(lines))
		// fmt.Println(i, w)
		start := encodeLines(lines)
		if first, ok := firstSeen[start]; ok {
			period := i - first
			if (times-first)%period == 0 {
				return lines
			}
		}
		if spinOut, ok := cache[start]; ok {
			lines = copyLines(spinOut)
			continue
		}
		lines = tiltUp(lines)
		lines = tiltLeft(lines)
		lines = tiltDown(lines)
		lines = tiltRight(lines)
		cache[start] = copyLines(lines)
		firstSeen[start] = i
	}
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

func tiltDown(lines []string) []string {
	slices.Reverse(lines)
	tiltUp(lines)
	slices.Reverse(lines)
	return lines
}

func tiltLeft(lines []string) []string {
	for i := 0; i < len(lines); i += 1 {
		for j, next := 0, 0; j < len(lines[i]); j += 1 {
			switch lines[i][j] {
			case '.':
				_ = 0
			case 'O':
				if next != j {
					lines[i] = lines[i][:next] + "O" + lines[i][next+1:]
					lines[i] = lines[i][:j] + "." + lines[i][j+1:]
				}
				next += 1
			case '#':
				next = j + 1
			}
		}
	}
	return lines
}

func reverseString(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i -= 1 {
		result += string(s[i])
	}
	return result
}

func tiltRight(lines []string) []string {
	for i := 0; i < len(lines); i += 1 {
		lines[i] = reverseString(lines[i])
	}
	tiltLeft(lines)
	for i := 0; i < len(lines); i += 1 {
		lines[i] = reverseString(lines[i])
	}
	return lines
}
