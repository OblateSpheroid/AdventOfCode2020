package main

import (
	"fmt"
)

type Line struct {
	lower  int
	upper  int
	letter byte
	pw     string
}

func isValid(l Line) bool {
	i := 0
	for _, v := range l.pw {
		if v == rune(l.letter) {
			i++
		}
	}
	if i >= l.lower && i <= l.upper {
		return true
	}
	return false
}

func isNewValid(l Line) bool {
	if l.pw[l.lower-1] == l.letter && l.pw[l.upper-1] == l.letter {
		return false
	}
	if l.pw[l.lower-1] == l.letter || l.pw[l.upper-1] == l.letter {
		return true
	}
	return false
}

func main() {
	i := 0
	j := 0
	lines := parseFile("data.txt")
	for _, line := range lines {
		if isValid(line) {
			i++
		}
		if isNewValid(line) {
			j++
		}
	}
	fmt.Printf("Answer 1: %v\n", i)
	fmt.Printf("Answer 2: %v\n", j)
}
