package main

import (
	"testing"
)

var data = parseFile("test1.txt")
var sol1 = checkLoop(data, 5)

func TestPart1(t *testing.T) {
	actual := sol1
	expected := 127
	if expected != actual {
		t.Errorf("Part 1 is wrong: expected %d, got %d\n", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	low, high := findContigSet(sol1, data)
	actual := low + high
	expected := 62
	if expected != actual {
		t.Errorf("Part 1 is wrong: expected %d, got %d\n", expected, actual)
	}
}
