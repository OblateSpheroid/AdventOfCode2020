package main

import (
	"fmt"
)

func playGame(start []int, n int) int {
	// given starting sequence, what is the nth number in game?
	m := make(map[int]int) // map number spoken : last last time it was spoken
	i := 1                 // counter
	last := 0              // last number spoken
	// run through starting sequence
	for _, w := range start {
		if i != 1 {
			m[last] = i - 1
		}
		last = w
		i++
	}

	// run through game until nth number has been spoken
	w := 0 // init var
	for i <= n {
		if m[last] == 0 {
			w = 0 // speak 0 if has not been spoken before
		} else {
			w = (i - 1) - m[last] // speak how many turns since last spoken
		}
		m[last] = i - 1 // update last time previous word spoken
		last = w
		i++
	}
	return last
}

func main() {
	test := []int{0, 3, 6}
	fmt.Println(playGame(test, 10) == 0)

	data := []int{17, 1, 3, 16, 19, 0}
	fmt.Printf("Answer 1: %d\n", playGame(data, 2020))
	/* I suspect there's a more efficient way to find nth number
	without going through each iteration */
	fmt.Printf("Answer 2: %d\n", playGame(data, 30000000))
}
