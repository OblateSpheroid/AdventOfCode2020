package main

import (
	"fmt"
	"sort"
)

func copySlice(sl []int) []int {
	tmp := make([]int, len(sl))
	for i, v := range sl {
		tmp[i] = v
	}
	return tmp
}

func sortSlice(sl []int) []int {
	sorted := sort.IntSlice(copySlice(sl))
	sort.Sort(sorted)
	return sorted
}

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

func sameSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func hasTried(s []int, tried *[][]int) bool {
	for _, v := range *tried {
		if sameSlice(s, v) {
			return true
		}
	}
	return false
}

func countArrange(sorted []int, tried *[][]int) int {
	if hasTried(sorted, tried) {
		return 0 // this combination has been tried already
	}
	*tried = append(*tried, sorted) // add current set to tried list
	c := 0
	dispose := []int{}
	for i, v := range sorted {
		if i == 0 {
			if v <= 1 && (sorted[1]-v) <= 1 {
				c++
				dispose = append(dispose, i)
			}
			continue
		}
		if i == (len(sorted) - 1) {
			continue // can never remove last adapter, will always be needed
		}
		if (v-sorted[i-1]) <= 1 && (sorted[i+1]-v) <= 1 {
			c++
			dispose = append(dispose, i)
		}
	}

	if len(dispose) > 0 {
		for _, v := range dispose {
			tmp := append(copySlice(sorted[:v]), sorted[v+1:]...) // remove an unneccessary adapter
			c += countArrange(tmp, tried)                         // recurse
		}
	}
	return c
}

func main() {
	test := parseFile("test1.txt")
	test = sortSlice(test)
	test1, test3 := countDiffs(test)
	tried := [][]int{}
	fmt.Printf("test answer 1: %d\n", test1*test3)
	fmt.Printf("test answer 2: %d\n", countArrange(test, &tried))

	data := parseFile("data.txt")
	sorted := sortSlice(data)
	diff1, diff3 := countDiffs(sorted)
	fmt.Printf("Answer 1: %d\n", diff1*diff3)
}
