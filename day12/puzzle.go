package main

import (
	"fmt"
)

var m = map[rune]Point{
	// map ordinal directions to x,y weights
	'N': {0, 1},
	'S': {0, -1},
	'E': {1, 0},
	'W': {-1, 0},
}

type Instruction struct {
	order rune
	num   float64
}

type Ship struct {
	pt  Point   // current position
	deg float64 // current direction in degrees
}

type Waypoint struct {
	pt    Point
	deg   float64
	shpPt Point // pt of ship
}

func (s *Ship) move(inst Instruction) {
	if inst.order == 'R' {
		s.deg -= inst.num // turning right decreases degrees
	} else if inst.order == 'L' {
		s.deg += inst.num // turning left increases degrees
	} else if inst.order == 'F' {
		s.pt = s.pt.fromDeg(s.deg, float64(inst.num))
	} else {
		s.pt = s.pt.add(m[inst.order].magnify(inst.num))
	}
}

func (w *Waypoint) move(inst Instruction) {
	if inst.order == 'R' {
		dist := w.pt.euclid(w.shpPt)
		w.deg -= inst.num // rotate right decreases degrees
		w.pt = w.shpPt.fromDeg(w.deg, dist)
	} else if inst.order == 'L' {
		dist := w.pt.euclid(w.shpPt)
		w.deg -= inst.num // rotate left decreases degrees
		w.pt = w.shpPt.fromDeg(w.deg, dist)
	} else if inst.order == 'F' {
		dist := float64(inst.num) * w.pt.euclid(w.shpPt)
		w.shpPt = w.shpPt.fromDeg(w.deg, dist)
		w.pt = w.pt.fromDeg(w.deg, dist)
	} else {
		w.pt = w.pt.add(m[inst.order].magnify(inst.num))
		w.deg = w.pt.toDeg(w.shpPt) // recalculate degrees when waypoint moves
	}
}

func loop(s *Ship, inst []Instruction) float64 {
	for _, i := range inst {
		s.move(i)
	}
	md := s.pt.dist(Point{0, 0}) // Manhattan distance from origin
	return md
}

func loop2(w *Waypoint, inst []Instruction) float64 {
	for _, i := range inst {
		w.move(i)
	}
	md := w.shpPt.dist(Point{0, 0})
	return md
}

func main() {
	test := parseFile("test1.txt")
	t := Ship{deg: 0}
	fmt.Println(25 == loop(&t, test))

	t2 := Waypoint{pt: Point{10, 1}, shpPt: Point{0, 0}}
	t2.deg = t2.pt.toDeg(Point{})
	fmt.Println(286 == loop2(&t2, test))

	data := parseFile("data.txt")
	s := Ship{deg: 0} // start facing east
	fmt.Printf("Answer 1: %v\n", loop(&s, data))

	data2 := parseFile("data.txt")
	w := Waypoint{pt: Point{10, 1}, shpPt: Point{0, 0}}
	w.deg = w.pt.toDeg(Point{})
	fmt.Printf("Answer 2: %v\n", loop2(&w, data2))
}
