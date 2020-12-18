package main

import "fmt"

func run(s Seats) (Seats, bool) {
	// apply rules to grid once and output new grid, and boolean if grid is changed
	news := s.makeCopy() // fill new grid without modifying old one
	for i, v := range s {
		for j := range v {
			if news[i][j].seat && !news[i][j].occupied && s.lookAround(j, i) == 0 {
				news[i][j].occupied = true // becomes occupied
				continue
			}
			if news[i][j].seat && news[i][j].occupied && s.lookAround(j, i) >= 4 {
				news[i][j].occupied = false // seat empties
			}
		}
	}
	return news, s.same(news)
}

func run2(s Seats) (Seats, bool) {
	// apply rules part 2
	news := s.makeCopy() // fill new grid without modifying old one
	for i, v := range s {
		for j := range v {
			if news[i][j].seat && !news[i][j].occupied && s.lookAround2(j, i) == 0 {
				news[i][j].occupied = true // becomes occupied
				continue
			}
			if news[i][j].seat && news[i][j].occupied && s.lookAround2(j, i) >= 5 {
				news[i][j].occupied = false // seat empties
			}
		}
	}
	return news, s.same(news)
}

func loop(s Seats, f func(s Seats) (Seats, bool)) int {
	// loop rules until steady-state, then return occupied count
	done := false
	for !done {
		s, done = f(s) // loop until not changes
	}
	return s.countOccupied()
}

func main() {
	test := parseFile("test1.txt")
	fmt.Println(loop(test, run) == 37)
	fmt.Println(loop(test, run2) == 26)

	data := parseFile("data.txt")
	fmt.Printf("Answer 1: %d\n", loop(data, run))
	fmt.Printf("Answer 2: %d\n", loop(data, run2))
}
