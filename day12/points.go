package main

import "math"

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

func (p Point) magnify(i float64) Point {
	n := Point{}
	n.x = p.x * float64(i)
	n.y = p.y * float64(i)
	return n
}

func (a Point) dist(b Point) float64 {
	// retrun Manhattan distance to another point
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func (a Point) euclid(b Point) float64 {
	// return Euclidian distance between two point
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

func (p Point) toDeg(o Point) float64 {
	// calculate degree of angle btwn origin o and point p
	radian := math.Atan((p.y - o.y) / (p.x - o.x))
	return radian * 180 / math.Pi
}

func (p Point) fromDeg(deg float64, dist float64) Point {
	// from starting point p, use theta and distance to find new point
	radian := float64(deg) * math.Pi / 180.0
	n := Point{dist * math.Cos(radian), dist * math.Sin(radian)}
	return p.add(n)
}
