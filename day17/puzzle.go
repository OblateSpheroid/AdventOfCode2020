package main

import (
	"fmt"
	"strings"
)

var data = `####.#..
.......#
#..#####
.....##.
##...###
#..#.#.#
.##...#.
#...##..`

var test = `.#.
..#
###`

type Point struct {
	x, y, z, w int
}

func (p Point) getNeighbors(d int) []Point {
	// return slice of all neighbors, depending on number of dimensions
	ns := []Point{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if d == 3 {
					if i != 0 || j != 0 || k != 0 {
						ns = append(ns, Point{p.x + i, p.y + j, p.z + k, 0}) // w always 0 in 3D
					}
				} else if d == 4 {
					for l := -1; l <= 1; l++ {
						if i != 0 || j != 0 || k != 0 || l != 0 {
							ns = append(ns, Point{p.x + i, p.y + j, p.z + k, p.w + l})
						}
					}
				}
			}
		}
	}
	return ns
}

func (p Point) countActiveNeighbors(g Grid, d int) int {
	// get number of active neighbors, depending on dimensions
	c := 0
	ns := p.getNeighbors(d)
	for _, n := range ns {
		if g[n] {
			c++
		}
	}
	return c
}

type Grid map[Point]bool // keeps track of which points are active

func (g Grid) update(d int) Grid {
	// make points active or unactive based on neighbors
	newg := make(Grid)
	for p := range g { // current grid
		ans := p.countActiveNeighbors(g, d)
		if g[p] && ans != 2 && ans != 3 {
			newg[p] = false
		} else if !g[p] && ans == 3 {
			newg[p] = true
		} else {
			newg[p] = g[p]
		}
		for _, n := range p.getNeighbors(d) { // same rules for outside neighbors
			ans2 := n.countActiveNeighbors(g, d)
			if !g[n] && ans2 == 3 {
				newg[n] = true
			}
		}
	}
	return newg
}

func (g Grid) count() int {
	// count total active points in grid
	c := 0
	for _, v := range g {
		if v {
			c++
		}
	}
	return c
}

func (g Grid) boot(cycles, d int) Grid {
	// update certain number of times (cycles)
	n := g.update(d) // first cycle
	for i := 2; i <= cycles; i++ {
		n = n.update(d)
	}
	return n
}

func initGrid(s string) Grid {
	// initialize grid from input data
	g := make(Grid)
	lines := strings.Split(s, "\n")
	for y, line := range lines {
		for x, c := range line {
			p := Point{x, y, 0, 0} // input data always 2D
			if c == '#' {
				g[p] = true
			} else {
				g[p] = false
			}
		}
	}
	return g
}

func main() {
	t := initGrid(test)
	t = t.boot(6, 3)
	fmt.Println(t.count() == 112)

	g := initGrid(data)
	g = g.boot(6, 3)
	fmt.Printf("Answer 1: %d\n", g.count())

	// Part 2: 4D
	t2 := initGrid(test)
	t2 = t2.boot(6, 4)
	fmt.Println(t2.count() == 848)

	g2 := initGrid(data)
	g2 = g2.boot(6, 4)
	fmt.Printf("Answer 2: %d\n", g2.count())
}
