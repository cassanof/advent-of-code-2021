package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/elleven11/advent-of-code-2021/utils"
	"github.com/thoas/go-funk"
)

// https://adventofcode.com/2021/day/2

func main() {
	lines, err := utils.ReadLines("./day2/input")
	utils.CheckPanic(err)

	commands := funk.Map(lines, func(s string) string {
		return s[:strings.IndexByte(s, ' ')]
	}).([]string)
	amounts := funk.Map(lines, func(s string) int {
		amt, err := strconv.Atoi(s[strings.IndexByte(s, ' ')+1:])
		utils.CheckPanic(err)
		return amt
	}).([]int)

	fmt.Println(calcPosition(commands, amounts))
}

func calcPosition(cmds []string, amts []int) int {
	h := 0
	d := 0
	aim := 0

	for i, cmd := range cmds {
		switch cmd {
		case "forward":
			amt := amts[i]
			h += amt
			d += aim * amt
		case "up":
			aim -= amts[i]
		case "down":
			aim += amts[i]
		}
	}
	return h * d
}
