package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func part2(input string) int {

	r1, _ := regexp.Compile("one|two|three|four|five|six|seven|eight|nine|[0-9]")
	r2, _ := regexp.Compile("eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|[0-9]")

	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		first := r1.FindString(line)
		last := reverse(r2.FindString(reverse(line)))

		val, _ := strconv.Atoi(fmt.Sprintf("%d%d", translate(first), translate(last)))

		sum += val
	}

	return sum
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func translate(s string) int {
	switch s {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		v, _ := strconv.Atoi(s)
		return v
	}
}
