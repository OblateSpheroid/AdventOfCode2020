package main

import (
	"fmt"
)

func find2(a []int) (int, int) {
	num := len(a)
	for i, x := range a {
		for _, y := range a[i+1 : num] {
			if (x + y) == 2020 {
				return x, y
			}
		}
	}
	return 0, 0
}

func find3(a []int) (int, int, int) {
	num := len(a)
	for i, x := range a[:num-2] {
		for j, y := range a[i+1 : num-1] {
			for _, z := range a[j+1 : num] {
				if (x + y + z) == 2020 {
					return x, y, z
				}
			}
		}
	}
	return 0, 0, 0
}

func main() {
	a := parseFile("data.txt")
	x, y := find2(a)
	fmt.Printf("Answer 1: %v\n", x*y)
	t, u, v := find3(a)
	fmt.Printf("Answer 2: %v\n", t*u*v)
}
