package main

import (
	"fmt"
)

type Rules map[string]map[string]int

func isIn(key string, m map[string]int) bool {
	// test if string is a key in a map
	for k := range m {
		if k == key {
			return true
		}
	}
	return false
}

func mapToSlice(m map[string]int) []string {
	// convert map to a slice of unique strings
	s := []string{}
	for k := range m {
		s = append(s, k)
	}
	return s
}

func findRecursive(r Rules, m map[string]int) []string {
	// recursively search for bags that contain an initial set of bags
	len1 := len(m)
	for k, v := range r {
		for bag := range m {
			if isIn(bag, v) {
				m[k] = 1
			}
		}
	}
	if len(m) > len1 {
		_ = findRecursive(r, m)
	}
	return mapToSlice(m)
}

func main() {
	rules := parseFile("data.txt")
	fmt.Println(rules["vibrant violet"])
	init := map[string]int{"shiny gold": 1} // initial map with "shiny gold"
	sl := findRecursive(rules, init)
	fmt.Printf("Answer 1: %d\n", len(sl)-1) // but "shiny gold" cannot contain itself
}
