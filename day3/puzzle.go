package main

import (
	"fmt"
)

type Pos struct {
	// describes a position on the grid as row, column
	row int
	col int
}

type Moves struct {
	// describes movement, as number of right moves, number of down moves
	right int
	down  int
}

func moveDown(p Pos, mv Moves, l int) Pos {
	// takes a starting position and moves to return ending position
	var new Pos
	new.row = p.row + mv.down
	if (p.col + mv.right) >= l { // "wrap around"
		new.col = (p.col + mv.right) - l
	} else { // normal shift
		new.col = p.col + mv.right
	}
	return new
}

func isTree(grid []string, p Pos) bool {
	// determine if a given position is a tree
	return grid[p.row][p.col] == 35
}

func countTrees(lines []string, mv Moves) int {
	// given a grid and movement strategy, returns
	// number of trees on route
	pos := Pos{0, 0}             // start at (0, 0)
	tot := 0                     // initialize total count of trees
	line_length := len(lines[0]) // should all be same length
	for i := 0; i < len(lines); i += mv.down {
		if isTree(lines, pos) {
			tot++
		}
		pos = moveDown(pos, mv, line_length)
	}
	return tot
}

func main() {
	lines := parseFile("data.txt")
	mv := Moves{right: 3, down: 1}
	fmt.Printf("Answer 1: %d\n", countTrees(lines, mv))

	p := 1 // initialize product
	m := map[Moves]int{
		{right: 1, down: 1}: 0,
		{right: 3, down: 1}: 0,
		{right: 5, down: 1}: 0,
		{right: 7, down: 1}: 0,
		{right: 1, down: 2}: 0,
	}
	for k := range m {
		m[k] = countTrees(lines, k)
		p *= m[k]
	}
	fmt.Println(m)
	fmt.Printf("Answer 2: %d\n", p)
}
