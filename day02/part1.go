package main

import (
	"strconv"
	"strings"
)

func part1(record string, contents struct{ red, green, blue int }) int {
	lines := strings.Split(record, "\n")
	totalCount := 0
GameLoop:
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		game := parts[0]
		rounds := strings.Split(parts[1], "; ")

		gameIdStr := strings.Split(game, " ")[1]
		gameId, _ := strconv.Atoi(gameIdStr)

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
					if contents.red < count {
						continue GameLoop
					}
				case "green":
					if contents.green < count {
						continue GameLoop
					}
				case "blue":
					if contents.blue < count {
						continue GameLoop
					}
				}
			}

		}
		totalCount += gameId
	}

	return totalCount
}
