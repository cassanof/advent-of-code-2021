package main

// https://adventofcode.com/2021/day/1

import (
	"fmt"

	"github.com/elleven11/advent-of-code-2021/utils"
)

func main() {
	lines, err := utils.ReadLinesAtoi("./day1/input")
	utils.CheckPanic(err)

	fmt.Println(countIncreses(lines))
}

func countIncreses(arr []int) int {
	count := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			count += 1
		}
	}

	return count
}
