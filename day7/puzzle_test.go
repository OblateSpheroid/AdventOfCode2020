package main

import (
	"testing"
)

var rules1 = parseFile("test1.txt")
var rules2 = parseFile("test2.txt")

func TestPart1(t *testing.T) {
	init := map[string]int{"shiny gold": 1} // initial map with "shiny gold"
	sol := findRecursive(rules1, init)
	actual := len(sol) - 1 // but "shiny gold" cannot contain itself
	expected := 4
	if expected != actual {
		t.Errorf("Part 1 is wrong: expected %d, got %d\n", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	sol1 := countRecursive(rules1, "shiny gold", 1)
	sol2 := countRecursive(rules2, "shiny gold", 1)
	expected1 := 32
	expected2 := 126
	if expected1 != sol1 {
		t.Errorf("Part 2 is wrong, expected %d, got %d\n", expected1, sol1)
	}
	if expected2 != sol2 {
		t.Errorf("Part 2 is wrong, expected %d, got %d\n", expected2, sol2)
	}
}
