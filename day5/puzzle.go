package main

import (
	"fmt"
	"math"
)

type Seat struct {
	row int
	col int
	id  int
}

func isIn(i int, is []int) bool {
	for _, v := range is {
		if i == v {
			return true
		}
	}
	return false
}

func parseSeat(s string) Seat {
	var seat Seat
	for i, v := range s[:7] {
		if v == 'B' {
			seat.row += int(math.Pow(2, float64(6-i)))
		}
	}
	for i, v := range s[7:] {
		if v == 'R' {
			seat.col += int(math.Pow(2, float64(2-i)))
		}
	}
	seat.id = (seat.row * 8) + seat.col
	return seat
}

func findHighest(s []string) int {
	highest := 0
	for _, v := range s {
		seat := parseSeat(v)
		if seat.id > highest {
			highest = seat.id
		}
	}
	return highest
}

func findMySeat(s []string, highest int) int {
	all_ids := []int{}
	for _, v := range s {
		all_ids = append(all_ids, parseSeat(v).id)
	}
	for i := 0; i <= highest; i++ {
		if !isIn(i, all_ids) && isIn(i-1, all_ids) && isIn(i+1, all_ids) {
			return i
		}
	}
	return 0
}

func main() {
	test_cases := []string{"BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}
	for _, v := range test_cases {
		fmt.Printf("%+v\n", parseSeat(v))
	}
	s := parseFile("data.txt")
	highest := findHighest(s)
	fmt.Printf("Answer 1: %d\n", highest)
	fmt.Printf("Answer 2: %d\n", findMySeat(s, highest))
}
