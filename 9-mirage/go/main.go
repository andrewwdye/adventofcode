package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func solveFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return solve(file)
}

func solve(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums := lo.Map(strings.Split(line, " "), func(s string, _ int) int {
			return lo.Must(strconv.Atoi(s))
		})
		sum += prev(nums)
	}
	return sum, nil
}

func next(nums []int) int {
	if lo.EveryBy(nums, func(n int) bool {
		return n == 0
	}) {
		return 0
	}
	diffs := make([]int, len(nums)-1)
	for i := range diffs {
		diffs[i] = nums[i+1] - nums[i]
	}
	return nums[len(nums)-1] + next(diffs)
}

func prev(nums []int) int {
	if lo.EveryBy(nums, func(n int) bool {
		return n == 0
	}) {
		return 0
	}
	diffs := make([]int, len(nums)-1)
	for i := range diffs {
		diffs[i] = nums[i+1] - nums[i]
	}
	return nums[0] - prev(diffs)
}

func main() {
	val, err := solveFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(val)
}
