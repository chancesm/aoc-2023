package main

import (
	"math"
	"slices"
	"strings"
)

func part1(input string) int {
	lines := strings.Split(input, "\n")
	totalCount := float64(0)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		gameParts := strings.Split(parts[1], " | ")

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
			totalCount += math.Pow(2, float64(matches-1))
		}
	}

	return int(totalCount)
}
