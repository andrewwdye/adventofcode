package main

import (
	"fmt"
	"os"
)

func solve(times, distances []int) (int, error) {
	if len(times) != len(distances) {
		return 0, fmt.Errorf("invalid input")
	}

	prod := 1
	for i := range times {
		prod *= waysToWin(times[i], distances[i])
	}

	fmt.Println(prod)
	return prod, nil
}

func waysToWin(time, distance int) int {
	ways := 0
	for speed := 0; speed < time; speed++ {
		if speed*(time-speed) > distance {
			ways += 1
		}
	}
	return ways
}

func main() {
	_, err := solve(
		// []int{54, 94, 65, 92},
		// []int{302, 1476, 1029, 1404},
		[]int{54946592},
		[]int{302147610291404},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
