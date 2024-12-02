package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"strconv"

	"github.com/andrewwdye/adventofcode2023/7-camel-cards/go/cards"
)

func sumHandsFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return sumHands(file)
}

func sumHands(reader io.Reader) (int, error) {
	hands, err := getHands(reader)
	if err != nil {
		return 0, err
	}

	slices.SortFunc(hands, func(a, b cards.Hand) int {
		return a.Cmp(b)
	})

	sum := 0
	for i, hand := range hands {
		fmt.Printf("%04d: %s --> %d\n", i+1, hand, hand.Bet*(i+1))
		sum += hand.Bet * (i + 1)
	}

	fmt.Println(sum)
	return sum, nil
}

func getHands(reader io.Reader) ([]cards.Hand, error) {
	scanner := bufio.NewScanner(reader)
	re := regexp.MustCompile(`([0-9AKQJT]{5}) (\d+)`)
	hands := []cards.Hand{}
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) != 3 {
			return nil, fmt.Errorf("invalid input: %s", line)
		}
		hand := matches[1]
		bet, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, fmt.Errorf("invalid input: %s", line)
		}
		hands = append(hands, cards.NewHand(hand, bet, true))
	}
	return hands, nil
}

func main() {
	if _, err := sumHandsFile(os.Args[1]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
