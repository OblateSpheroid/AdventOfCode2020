package main

/* Create Point type with helper methods */

func abs(x float64) float64 {
	// return absolute value of a number
	if x < 0 {
		return -1 * x
	}
	return x
}

type Point struct {
	// a point on the grid
	x float64
	y float64
}

func (a Point) add(b Point) Point {
	// add a point to another point
	p := Point{}
	p.x = a.x + b.x
	p.y = a.y + b.y
	return p
}

func (p Point) magnify(i int) Point {
	n := Point{}
	n.x = p.x * float64(i)
	n.y = p.y * float64(i)
	return n
}

func (a Point) dist(b Point) float64 {
	// retrun Manhattan distance to another point
	return abs(a.x-b.x) + abs(a.y-b.y)
}
