package main

import (
	"aoc2020/helpers"
	"fmt"
)

type Rules map[string]map[string]int

func findRecursive(r Rules, m map[string]int) []string {
	// recursively search for bags that contain an initial set of bags
	len1 := len(m)
	for k, v := range r {
		for bag := range m {
			if helpers.IsInMap(bag, v) {
				m[k] = 1
			}
		}
	}
	if len(m) > len1 {
		_ = findRecursive(r, m)
	}
	return helpers.MapToSlice(m)
}

func countRecursive(r Rules, b string, p int) int {
	// r : total rule set
	// b : bag name, e.g., "shiny gold"
	// p : parent count
	i := 0
	bags := r[b]
	for bag, c := range bags {
		if c == 0 {
			continue // end of line
		}
		i += p * c // parents bags * bags in each parent
		i += countRecursive(r, bag, p*c)
	}
	return i
}

func main() {
	rules := parseFile("data.txt")
	init := map[string]int{"shiny gold": 1} // initial map with "shiny gold"
	sol1 := findRecursive(rules, init)
	fmt.Printf("Answer 1: %d\n", len(sol1)-1) // but "shiny gold" cannot contain itself

	sol2 := countRecursive(rules, "shiny gold", 1)
	fmt.Printf("Answer 2: %d\n", sol2)
}
