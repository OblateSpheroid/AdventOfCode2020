package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	lower  int
	upper  int
	letter byte
	pw     string
}

func parseFile(s string) []Line {
	f, _ := os.Open(s)
	defer f.Close()

	a := []Line{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var l Line
		tmp := strings.Split(scanner.Text(), " ")
		nums := strings.Split(tmp[0], "-")
		l.lower, _ = strconv.Atoi(nums[0])
		l.upper, _ = strconv.Atoi(nums[1])
		l.letter = tmp[1][0]
		l.pw = tmp[2]
		a = append(a, l)
	}

	return a
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
