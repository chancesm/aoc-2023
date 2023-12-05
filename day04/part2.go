package main

import (
	"slices"
	"strings"
)

func part2(input string) int {
	lines := strings.Split(input, "\n")
	cardLines := make([][]string, len(lines))
	cardCounts := make([]int, len(lines))
	totalCount := 0

	for i, line := range lines {
		cardLines[i] = strings.Split(line, ": ")
		cardCounts[i] = 1
	}
	for j, cardLine := range cardLines {
		// card := cardLine[0]
		line := cardLine[1]
		gameParts := strings.Split(line, " | ")

		winners := []string{}
		for _, numStr := range strings.Split(gameParts[0], " ") {
			if numStr == "" {
				continue
			}
			winners = append(winners, numStr)

		}
		matches := 0
		for _, numStr := range strings.Split(gameParts[1], " ") {
			if numStr == "" {
				continue
			}
			if slices.Contains(winners, numStr) {
				matches++
			}
		}
		if matches > 0 {
			for k := 1; k <= matches; k++ {
				for l := 1; l <= cardCounts[j]; l++{
					cardCounts[j+k] += 1
				}
			}
		}
	}
	for _, cardCount := range cardCounts {
		totalCount += cardCount
	}

	return totalCount

}
