package main

import (
	"aoc2020/helpers"
	"fmt"
)

type Cmd struct {
	op  string
	arg int
}

var isIn = helpers.IsIn

func run(c []Cmd) (int, bool) {
	line_count := len(c)
	acc := 0          // initialze at 0
	hasRun := []int{} // list of line numbers that have run
	exitNormal := true
	i := 0
	for i < line_count {
		if isIn(i, hasRun) {
			exitNormal = false
			break
		}
		hasRun = append(hasRun, i)
		if c[i].op == "acc" {
			acc += c[i].arg
			i++
			continue
		}
		if c[i].op == "jmp" {
			i += c[i].arg
			continue
		}
		if c[i].op == "nop" {
			i++
			continue
		}
	}
	return acc, exitNormal
}

func copySlice(c []Cmd) []Cmd {
	tmp := make([]Cmd, 0, len(c))
	for _, v := range c {
		tmp = append(tmp, v)
	}
	return tmp
}

func loopCheck(c []Cmd) int {
	// loop through intructions, changing one at a time
	// until one program exits normally
	try := 0
	check := false
	for i, cmd := range c {
		if cmd.op == "jmp" {
			trial := copySlice(c) // don't want to change data in place
			trial[i].op = "nop"
			try, check = run(trial)
			if check {
				break
			}
		}
		if cmd.op == "nop" {
			trial := copySlice(c) // don't want to change data in place
			trial[i].op = "jmp"
			try, check = run(trial)
			if check {
				break
			}
		}
	}
	return try
}

func main() {
	c := parseFile("data.txt")
	sol1, check1 := run(c)
	fmt.Printf("Answer 1: %d, %v\n", sol1, check1)
	fmt.Printf("Answer 2: %d\n", loopCheck(c))
}
