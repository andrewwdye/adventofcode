package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func sumCalibrationValuesFile(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}

	return sumCalibrationValues(f)
}

func sumCalibrationValues(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	sum := 0
	for s.Scan() {
		value, err := parseLine2(s.Text())
		if err != nil {
			return 0, err
		}
		sum += value
	}

	fmt.Printf("%d\n", sum)
	return sum, nil
}

// func parseLine(line string) (int, error) {
// 	digits := ""
// 	for i := range line {
// 		if line[i] >= '0' && line[i] <= '9' {
// 			digits += string(line[i])
// 			break
// 		}
// 	}
// 	for i := range line {
// 		j := len(line) - i - 1
// 		if line[j] >= '0' && line[j] <= '9' {
// 			digits += string(line[j])
// 			break
// 		}
// 	}
// 	if len(digits) == 0 {
// 		return 0, nil
// 	}
// 	return strconv.Atoi(digits)
// }

func parseLine2(line string) (int, error) {
	tokens := tokenize(line)
	if len(tokens) == 0 {
		return 0, nil
	}
	digits := tokens[0] + tokens[len(tokens)-1]
	return strconv.Atoi(digits)
}

func tokenize(line string) []string {
	result := make([]string, 0)
	for i := 0; i < len(line); i += 1 {
		if line[i] >= '0' && line[i] <= '9' {
			result = append(result, string(line[i]))
		} else {
			strs := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
			digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
			for j := range strs {
				if len(line[i:]) >= len(strs[j]) && line[i:i+len(strs[j])] == strs[j] {
					result = append(result, digits[j])
					break
				}
			}
		}
	}
	return result
}

func main() {
	_, err := sumCalibrationValuesFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
