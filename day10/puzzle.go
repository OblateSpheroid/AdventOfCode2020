package main

import (
	"aoc2020/helpers"
	"fmt"
)

func countDiffs(sorted []int) (int, int) {
	diff := make(map[int]int)
	for i, v := range sorted {
		if i == 0 {
			diff[v]++ // difference between outlet and first adapter
			continue
		}
		diff[v-sorted[i-1]]++
	}
	diff[3]++ // difference between last adapter and device
	return diff[1], diff[3]
}

func countPaths(sorted []int) int {
	m := make(map[int]int)              // m[i] = number of paths to i jolts
	m[0] = 1                            // start at outlet
	last_adapt := sorted[len(sorted)-1] // max adapter is last in list
	device := last_adapt + 3            // jolts needed by device
	for i := 1; i <= last_adapt; i++ {
		if helpers.IsInSorted(i, sorted) {
			m[i] = m[i-1] + m[i-2] + m[i-3]
		}
	}
	m[device] = m[device-1] + m[device-2] + m[device-3]
	return m[device] // number of paths to get to device
}

func main() {
	test := parseFile("test1.txt")
	test = helpers.SortSlice(test)
	test1, test3 := countDiffs(test)
	fmt.Printf("test answer 1: %d\n", test1*test3)
	fmt.Printf("test answer 2: %d\n", countPaths(test))

	data := helpers.SortSlice(parseFile("data.txt"))
	diff1, diff3 := countDiffs(data)
	fmt.Printf("Answer 1: %d\n", diff1*diff3)
	fmt.Printf("Answer 2: %d\n", countPaths(data))
}
