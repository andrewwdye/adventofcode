package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func countValidFile(file string, r, g, b int) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return countValid(f, r, g, b)
}

func countValid(reader io.Reader, r, g, b int) (int, error) {
	s := bufio.NewScanner(reader)
	sum := 0
	for s.Scan() {
		v, err := validLine(s.Text(), r, g, b)
		if err != nil {
			return 0, err
		}
		sum += v
	}

	fmt.Printf("%d\n", sum)
	return sum, nil
}

func validLine(line string, r, g, b int) (int, error) {
	game_re := regexp.MustCompile(`Game (\d+)`)
	red_re := regexp.MustCompile(`(\d+) red`)
	green_re := regexp.MustCompile(`(\d+) green`)
	blue_re := regexp.MustCompile(`(\d+) blue`)

	game := game_re.FindStringSubmatch(line)
	if len(game) == 0 {
		return 0, fmt.Errorf("missing game from line: %s", line)
	}

	checkColor := func(re *regexp.Regexp, max int) (bool, error) {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, m := range matches {
			count, err := strconv.Atoi(m[1])
			if err != nil {
				return false, err
			}
			if count > max {
				return false, nil
			}
		}
		return true, nil
	}
	if valid, err := checkColor(red_re, r); !valid {
		return 0, err
	}
	if valid, err := checkColor(green_re, g); !valid {
		return 0, err
	}
	if valid, err := checkColor(blue_re, b); !valid {
		return 0, err
	}

	return strconv.Atoi(game[1])
}

func sumPowerFile(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}

	return sumPower(f)
}

func sumPower(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	sum := 0
	for s.Scan() {
		power, err := power(s.Text())
		if err != nil {
			return 0, err
		}
		sum += power
	}

	fmt.Printf("%d\n", sum)
	return sum, nil
}

func power(line string) (int, error) {
	red_re := regexp.MustCompile(`(\d+) red`)
	green_re := regexp.MustCompile(`(\d+) green`)
	blue_re := regexp.MustCompile(`(\d+) blue`)

	result := 1
	for _, re := range []*regexp.Regexp{red_re, green_re, blue_re} {
		matches := re.FindAllStringSubmatch(line, -1)
		max := 0
		for _, m := range matches {
			count, err := strconv.Atoi(m[1])
			if err != nil {
				return 0, err
			}
			if count > max {
				max = count
			}
		}
		result *= max
	}
	return result, nil
}

func main() {
	// _, err := countValidFile(os.Args[1], 12, 13, 14)
	_, err := sumPowerFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
