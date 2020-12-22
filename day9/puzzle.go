package main

import (
	"aoc2020/helpers"
	"fmt"
)

func check(n int, sl []int) bool {
	// sl is the slice of the preamble
	// check that any two integers in preamble sum to n
	for i, v := range sl[:len(sl)-1] {
		for _, w := range sl[i+1:] {
			if (v + w) == n {
				return true
			}
		}
	}
	return false
}

func checkLoop(sl []int, pre int) int {
	// find first item in slice that is invalid
	ans := 0
	for i, v := range sl {
		if i < pre {
			continue // skip preamble
		}
		if check(v, sl[i-pre:i]) {
			continue // checks out
		} else {
			return v // v is first that don't check out
		}
	}
	return ans
}

func findContigSet(n int, sl []int) (int, int) {
	// return smallest and largest number of contiguous set
	for i := 0; i < len(sl)-1; i++ {
		for j := i + 1; j <= len(sl); j++ {
			if helpers.Sum(sl[i:j]) == n {
				return helpers.Min(sl[i:j]), helpers.Max(sl[i:j]) // found min and max numbers of set
			} else if helpers.Sum(sl[i:j]) > n {
				break // overshot, go to next i
			}
		}
	}
	return -1, -1 // to indicate no answer found
}

func main() {
	sl := parseFile("data.txt")
	sol1 := checkLoop(sl, 25)
	fmt.Printf("Answer 1: %d\n", sol1)
	low, high := findContigSet(sol1, sl)
	fmt.Printf("Answer 2: %d + %d = %d\n", low, high, low+high)
}
