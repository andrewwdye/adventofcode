package pkg

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Solve(reader io.Reader, diffs int) (int, error) {
	scanner := bufio.NewScanner(reader)
	sum := 0
	boards := [][]string{{}}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			boards = append(boards, []string{})
		} else {
			boards[len(boards)-1] = append(boards[len(boards)-1], line)
		}
	}
	for i, b := range boards {
		left, right := findReflection(b, diffs)
		if left == 0 && right == 0 {
			fmt.Println(i)
			fmt.Println(strings.Join(b, "\n"))
		}
		sum += left + right*100
	}
	return sum, nil
}

func findReflection(lines []string, diffs int) (int, int) {
	for i := 1; i < len(lines[0]); i++ {
		if checkColumn(lines, i) == diffs {
			return i, 0
		}
	}
	for i := 1; i < len(lines); i++ {
		if checkRow(lines, i) == diffs {
			return 0, i
		}
	}
	return 0, 0
}

func checkColumn(lines []string, column int) int {
	diffs := 0
	for left, right := column-1, column; left >= 0 && right < len(lines[0]); left, right = left-1, right+1 {
		for i := 0; i < len(lines); i++ {
			if lines[i][left] != lines[i][right] {
				diffs += 1
			}
		}
	}
	return diffs
}

func checkRow(lines []string, row int) int {
	diffs := 0
	for left, right := row-1, row; left >= 0 && right < len(lines); left, right = left-1, right+1 {
		for i := 0; i < len(lines[0]); i++ {
			if lines[left][i] != lines[right][i] {
				diffs += 1
			}
		}
	}
	return diffs
}
