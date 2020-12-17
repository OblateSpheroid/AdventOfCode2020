package main

import "fmt"

func findUnique(c []string) (string, string) {
	// given a card, return unique yes answers
	m := make(map[rune]int)
	union := []rune{}
	intersect := []rune{}
	people := 0
	for _, line := range c { // create a map of all runes used
		people++
		for _, char := range line {
			m[char]++
		}
	}
	for k, v := range m {
		union = append(union, k) // slice of all runes that appeard at least once
		if v == people {
			intersect = append(intersect, k) // slice of runes common to everyone
		}
	}
	return string(union), string(intersect) // convert rune slices to strings
}

func main() {
	cards := parseFile("data.txt")
	count1 := 0
	count2 := 0
	for _, card := range cards {
		union, intersect := findUnique(card)
		count1 += len(union)
		count2 += len(intersect)
	}
	fmt.Printf("Answer 1: %d\n", count1)
	fmt.Printf("Answer 2: %d\n", count2)
}
