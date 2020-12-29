package main

import (
	"aoc2020/helpers"
	"fmt"
)

func countInvalidValues(t []int, r map[string][]int) int {
	c := 0
	for _, n := range t {
		b := false
		for _, v := range r {
			if helpers.IsIn(n, v) {
				b = true
				break
			}
		}
		if !b {
			c += n
		}
	}
	return c
}

func errorRate(tkts [][]int, r map[string][]int) int {
	c := 0
	for _, t := range tkts {
		c += countInvalidValues(t, r)
	}
	return c
}

// Part 2
func checkValid(t []int, r map[string][]int) bool {
	for _, n := range t {
		b := false
		for _, v := range r {
			if helpers.IsIn(n, v) {
				b = true
				break
			}
		}
		if !b {
			return false
		}
	}
	return true
}

func goodList(tkts [][]int, r map[string][]int) [][]int {
	good := [][]int{}
	for _, t := range tkts {
		if checkValid(t, r) {
			good = append(good, t)
		}
	}
	return good
}

func findOrder(tkts [][]int, r map[string][]int) map[string][]int {
	n := len(tkts[0])           // length of each ticket should be same
	m := make(map[string][]int) // map of possibilities
	for k := range r {
		m[k] = helpers.MakeSeq(0, n-1) // populate possibilities
	}
	for _, t := range tkts {
		for i, v := range t {
			for k := range r {
				if !helpers.IsIn(v, r[k]) {
					m[k] = helpers.Drop(m[k], i) // process of elimination
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		m = whittle(m) // try to get everything down to 1 possibility
	}
	return m
}

func whittle(m map[string][]int) map[string][]int {
	for k, v := range m {
		if len(v) == 1 { // if only 1 possibility, no others can have it
			for k2 := range m {
				if k2 != k {
					m[k2] = helpers.Drop(m[k2], v[0])
				}
			}
		}
	}
	return m
}

func answer2(t []int, m map[string][]int) int {
	// take single ticket and map of possitions
	// return product of 6 fields that start with 'departure'
	x1 := t[m["departure location"][0]]
	x2 := t[m["departure station"][0]]
	x3 := t[m["departure platform"][0]]
	x4 := t[m["departure track"][0]]
	x5 := t[m["departure date"][0]]
	x6 := t[m["departure time"][0]]
	return x1 * x2 * x3 * x4 * x5 * x6
}

func main() {
	tickets := parseTickets("nearby.txt")
	myTicket := parseMyTicket("ticket.txt")
	rules := parseRules("rules.txt")
	fmt.Printf("Answer 1: %d\n", errorRate(tickets, rules))

	good := goodList(tickets, rules)
	order := findOrder(good, rules)
	fmt.Printf("Answer 2: %v\n", answer2(myTicket, order))
}
