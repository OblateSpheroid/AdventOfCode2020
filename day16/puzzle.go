package main

import (
	"aoc2020/helpers"
	"fmt"
)

// Rules: map of field names to valid values
// Poss: map of field names to possible index on the ticket
type Tkt []int
type Tkts []Tkt
type Rules map[string][]int
type Poss map[string][]int

func countInvalidValues(t Tkt, r Rules) int {
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

func errorRate(tkts Tkts, r Rules) int {
	c := 0
	for _, t := range tkts {
		c += countInvalidValues(t, r)
	}
	return c
}

// Part 2
func checkValid(t Tkt, r Rules) bool {
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

func goodList(tkts Tkts, r Rules) Tkts {
	good := Tkts{}
	for _, t := range tkts {
		if checkValid(t, r) {
			good = append(good, t)
		}
	}
	return good
}

func whittle(p Poss) Poss {
	for k, v := range p {
		if len(v) == 1 { // if only 1 possibility, no others can have it
			for k2 := range p {
				if k2 != k {
					p[k2] = helpers.Drop(p[k2], v[0])
				}
			}
		}
	}
	return p
}

func findOrder(tkts Tkts, r Rules) Poss {
	n := len(tkts[0]) // length of each ticket should be same
	p := make(Poss)   // map of possibilities
	for k := range r {
		p[k] = helpers.MakeSeq(0, n-1) // populate possibilities
	}
	// round 1: remove possibilities based on rule definitions
	for _, t := range tkts {
		for i, v := range t {
			for k := range r {
				if !helpers.IsIn(v, r[k]) {
					p[k] = helpers.Drop(p[k], i)
				}
			}
		}
	}
	// round 2: whittle down through process of elimination
	for i := 0; i < n; i++ { // will need to run at most n times
		p = whittle(p)
	}
	return p
}

func answer2(t Tkt, m Poss) int {
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
