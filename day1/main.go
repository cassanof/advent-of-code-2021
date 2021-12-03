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

	for i := 1; i < len(arr)-2; i++ {
		prev_win := arr[i] + arr[i+1] + arr[i+2]
		next_win := arr[i-1] + arr[i] + arr[i+1]
		if prev_win > next_win {
			count += 1
		}
	}

	return count
}
