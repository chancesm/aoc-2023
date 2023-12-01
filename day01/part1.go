package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part1(input string) int {
	lines := strings.Split(input, "\n")
	lineNums := make([]int, len(lines))
	for i, line := range lines {
		foundNums := make([]int, 0)
		for _, char := range line {
			if val, err := strconv.Atoi(string(char)); err == nil {
				foundNums = append(foundNums, val)
			}

		}
		v, _ := strconv.Atoi(fmt.Sprintf("%d%d", foundNums[0], foundNums[len(foundNums)-1]))
		lineNums[i] = v
	}
	sum := 0
	for _, val := range lineNums {
		sum += val
	}
	return sum
}
