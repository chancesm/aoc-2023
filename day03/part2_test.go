package main

import "testing"

func Test_part2(t *testing.T) {
	testInput := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	expected := 467835

	if actual := part2(testInput); actual != expected {
		t.Errorf("part1() = %v, expected %v", actual, expected)
	}
}
