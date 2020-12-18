package main

/* Create types for Seat and Seats with helper methods */

type Seat struct {
	seat     bool // false = floor
	occupied bool // false = empty
}

type Seats [][]Seat

func (s Seats) findSeat(x, y int) Seat {
	// helper to return Seat at point in grid
	if y >= len(s) || x >= len(s[0]) || x < 0 || y < 0 {
		return Seat{} // out of bounds = floor
	}
	return s[y][x]
}

func (ss *Seats) appendInPlace(s []Seat) {
	// helper to append a row to grid
	*ss = append(*ss, s)
}

func (s Seats) makeCopy() Seats {
	// make a copy of a grid
	seats := Seats{}
	row_len := len(s[0])
	for _, v := range s {
		row := make([]Seat, row_len)
		for j, w := range v {
			row[j] = w
		}
		seats.appendInPlace(row)
	}
	return seats
}

func (s Seats) lookAround(x, y int) int {
	// return count of occupied seats around a seat
	c := 0 // init counter
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue //don't look at seat itself
			}
			if s.findSeat(i, j).occupied {
				c++
			}
		}
	}
	return c
}

func (s Seats) lookAround2(x, y int) int {
	// return count of occupied seats, first chair in each direction
	c := 0 // init counter
	// loop through all directions
	row_len := len(s[0])
	row_num := len(s)
	type Point [2]int
	ds := [8]Point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, d := range ds {
		for a := [2]int{x + d[0], y + d[1]}; a[0] >= 0 && a[1] >= 0 && a[0] < row_len && a[1] < row_num; a = [2]int{a[0] + d[0], a[1] + d[1]} {
			if s.findSeat(a[0], a[1]).seat {
				if s.findSeat(a[0], a[1]).occupied {
					c++
				}
				break // don't keep looking after finding first seat
			}
		}
	}
	return c
}

func (s Seats) same(n Seats) bool {
	// determine if two grids are the same
	for i, v := range s {
		for j := range v {
			if s[i][j] != n[i][j] {
				return false // first discrepancy means not equal
			}
		}
	}
	return true
}

func (s Seats) countOccupied() int {
	// count total occupied seats in grid
	c := 0
	for i, v := range s {
		for j := range v {
			if s[i][j].occupied {
				c++
			}
		}
	}
	return c
}
