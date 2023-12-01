package main

import "testing"

func Test_part1(t *testing.T) {
	testInput := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`
	expectedOutput := 142
	if got := part1(testInput); got != expectedOutput {
		t.Errorf("part1() = %v, want %v", got, expectedOutput)
	}
}