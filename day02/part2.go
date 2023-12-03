package main

import (
	"strconv"
	"strings"
)

func part2(record string) int {
	lines := strings.Split(record, "\n")
	totalCount := 0
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		rounds := strings.Split(parts[1], "; ")

		maxRed := 0
		maxGreen := 0
		maxBlue := 0

		for _, round := range rounds {
			// round = "8 green, 6 blue, 20 red"
			// get color counts from the round
			colors := strings.Split(round, ", ")
			for _, colorCount := range colors {
				values := strings.Split(colorCount, " ")
				color := values[1]
				countStr := values[0]
				count, _ := strconv.Atoi(countStr)
				switch color {
				case "red":
					if count > maxRed {
						maxRed = count
					}
				case "green":
					if count > maxGreen {
						maxGreen = count
					}
				case "blue":
					if count > maxBlue {
						maxBlue = count
					}
				}
			}

		}
		totalCount += (maxRed * maxGreen * maxBlue)
	}

	return totalCount
}
