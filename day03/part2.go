package main

import (
	"regexp"
	"strconv"
	"strings"
)

type LineNums struct {
	nums    []string
	indices [][]int
}

func part2(input string) int {
	numberRe := regexp.MustCompile(`[0-9]+`)
	gearRe := regexp.MustCompile(`\*`)

	totalCount := 0
	lines := strings.Split(input, "\n")

	prevLineNums := LineNums{}
	currLineNums := LineNums{}
	nextLineNums := LineNums{}

	for i, line := range lines {

		if i == 0 {
			currLineNums.nums = numberRe.FindAllString(line, -1)
			currLineNums.indices = numberRe.FindAllStringIndex(line, -1)

			nextLineNums.nums = numberRe.FindAllString(lines[i+1], -1)
			nextLineNums.indices = numberRe.FindAllStringIndex(lines[i+1], -1)
		} else if i == len(lines)-1 {
			prevLineNums = currLineNums
			currLineNums = nextLineNums
			nextLineNums = LineNums{}
		} else {
			prevLineNums = currLineNums
			currLineNums = nextLineNums
			nextLineNums.nums = numberRe.FindAllString(lines[i+1], -1)
			nextLineNums.indices = numberRe.FindAllStringIndex(lines[i+1], -1)
		}

		gears := gearRe.FindAllStringIndex(line, -1)
		if len(gears) == 0 {
			continue
		}
		
		for _, gear := range gears {
			adjacentNums := []string{}
			location := gear[0]

			for k, lox1 := range prevLineNums.indices {
				if location >= lox1[0]-1 && location <= lox1[1] {
					adjacentNums = append(adjacentNums, prevLineNums.nums[k])
				}
			}
			for l, lox2 := range currLineNums.indices {
				if location >= lox2[0]-1 && location <= lox2[1] {
					adjacentNums = append(adjacentNums, currLineNums.nums[l])
				}
			}
			for m, lox3 := range nextLineNums.indices {
				if location >= lox3[0]-1 && location <= lox3[1] {
					adjacentNums = append(adjacentNums, nextLineNums.nums[m])
				}
			}
			if len(adjacentNums) == 2 {
				firstNum, _ := strconv.Atoi(adjacentNums[0])
				secondNum, _ := strconv.Atoi(adjacentNums[1])
				totalCount += (firstNum * secondNum)
			}
		}

	}

	return totalCount
}
