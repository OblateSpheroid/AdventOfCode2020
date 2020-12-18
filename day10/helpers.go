package main

// Some basic functions for working with slices

import (
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

func isIn(i int, sl []int) bool {
	// requires slice to be sorted
	for _, v := range sl {
		if i == v {
			return true
		}
		if v > i {
			return false // already passed target value
		}
	}
	return false
}
