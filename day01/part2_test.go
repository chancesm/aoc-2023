package main

import "testing"

func Test_part2(t *testing.T) {
	testInput := `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`
	expectedOutput := 281
	if got := part2(testInput); got != expectedOutput {
		t.Errorf("part1() = %v, want %v", got, expectedOutput)
	}
}
func Test_part2_tricks(t *testing.T) {
	testInput := `gnmkdm7sevenseven3four7fhrhppmtkpzvtlfqoneighth`
	expectedOutput := 78
	if got := part2(testInput); got != expectedOutput {
		t.Errorf("part1() = %v, want %v", got, expectedOutput)
	}
}
