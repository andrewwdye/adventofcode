package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	var result int
	switch os.Args[1] {
	case "1":
		result, err = solve1(f)
	default:
		err = fmt.Errorf("invalid argument")
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
}
