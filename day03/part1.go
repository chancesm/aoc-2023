package main

import (
	"regexp"
	"strconv"
	"strings"
)

func part1(input string) int {
	foundSymbols := struct{ before, current, after [][]int }{before: [][]int{}, current: [][]int{}, after: [][]int{}}
	symbolRe := regexp.MustCompile(`[^A-Za-z0-9\.]`)
	numberRe := regexp.MustCompile(`[0-9]+`)
	lines := strings.Split(input, "\n")

	finalCount := 0

	for i, line := range lines {
		lineNumbers := numberRe.FindAllString(line, -1)
		lineNumbersIn := numberRe.FindAllStringIndex(line, -1)

		if i == 0 {
			foundSymbols.current = symbolRe.FindAllStringIndex(line, -1)
			foundSymbols.after = symbolRe.FindAllStringIndex(lines[i+1], -1)
		} else if i == len(lines)-1 {
			foundSymbols.before = foundSymbols.current
			foundSymbols.current = foundSymbols.after
			foundSymbols.after = [][]int{}
		} else {
			foundSymbols.before = foundSymbols.current
			foundSymbols.current = foundSymbols.after
			foundSymbols.after = symbolRe.FindAllStringIndex(lines[i+1], -1)
		}
	NumberLoop:
		for j, number := range lineNumbers {
			firstCheck := lineNumbersIn[j][0] - 1
			lastCheck := lineNumbersIn[j][1]

			for _, symbol := range foundSymbols.before {
				if symbol[0] >= firstCheck && symbol[0] <= lastCheck {
					val, _ := strconv.Atoi(number)
					finalCount += val
					continue NumberLoop
				}
			}
			for _, symbol := range foundSymbols.current {
				if symbol[0] >= firstCheck && symbol[0] <= lastCheck {
					val, _ := strconv.Atoi(number)
					finalCount += val
					continue NumberLoop
				}
			}
			for _, symbol := range foundSymbols.after {
				if symbol[0] >= firstCheck && symbol[0] <= lastCheck {
					val, _ := strconv.Atoi(number)
					finalCount += val
					continue NumberLoop
				}
			}
		}

	}
	return finalCount
}
