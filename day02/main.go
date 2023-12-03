package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"

	"github.com/chancesm/aoc-2023/pkg/util"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 0, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", util.Ter(part == 0, "1 & 2", strconv.Itoa(part)))
	if part == 0 {
		runPart1()
		runPart2()
	} else if part == 1 {
		runPart1()
	} else if part == 2 {
		runPart2()
	}
}
func runPart1() {
	ans := part1(input, struct{red, green, blue int}{red: 12, green: 13, blue: 14})
	fmt.Println("Part 1 Output:", ans)
}
func runPart2() {
	ans := part2(input)
	fmt.Println("Part 2 Output:", ans)
}
