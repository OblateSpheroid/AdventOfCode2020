package main

import (
	"fmt"
	"math"
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
	num   int
}

type Ship struct {
	pt  Point // current position
	deg int   // current direction in degrees
}

func degToWeights(deg int) Point {
	// Determine x,y weights given a degree using trig
	radian := float64(deg) * math.Pi / 180.0
	return Point{math.Cos(radian), math.Sin(radian)}
}

func (s *Ship) move(inst Instruction) {
	if inst.order == 'R' {
		s.deg -= inst.num // turning right decreases degrees
	} else if inst.order == 'L' {
		s.deg += inst.num // turning left increases degrees
	} else if inst.order == 'F' {
		w := degToWeights(s.deg)
		s.pt = s.pt.add(w.magnify(inst.num)) // move forward
	} else {
		s.pt = s.pt.add(m[inst.order].magnify(inst.num))
	}
}

func loop(s *Ship, inst []Instruction) float64 {
	for _, i := range inst {
		s.move(i)
	}
	md := s.pt.dist(Point{0, 0}) // Manhattan distance from origin
	return math.Round(md)
}

func main() {
	test := parseFile("test1.txt")
	t := Ship{deg: 0}
	fmt.Println(25==loop(&t, test))

	data := parseFile("data.txt")
	s := Ship{deg: 0} // start facing east
	sol1 := loop(&s, data)
	fmt.Printf("Answer 1: %v\n", sol1)
}
