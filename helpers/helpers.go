package helpers

/* Some generic helper functions be used by different days */

import (
	"sort"
)

func CopySlice(sl []int) []int {
	tmp := make([]int, len(sl))
	for i, v := range sl {
		tmp[i] = v
	}
	return tmp
}

func SortSlice(sl []int) []int {
	sorted := sort.IntSlice(CopySlice(sl))
	sort.Sort(sorted)
	return sorted
}

func IsInSorted(i int, sl []int) bool {
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

func IsIn(i int, sl []int) bool {
	for _, v := range sl {
		if i == v {
			return true
		}
	}
	return false
}

func IsInMap(key string, m map[string]int) bool {
	// test if string is a key in a map
	for k := range m {
		if k == key {
			return true
		}
	}
	return false
}

func MapToSlice(m map[string]int) []string {
	// convert map to a slice of unique strings
	s := []string{}
	for k := range m {
		s = append(s, k)
	}
	return s
}

// Some basic math functions

func Sum(sl []int) int {
	// sum all integers in a slice
	tot := 0
	for _, v := range sl {
		tot += v
	}
	return tot
}

func Min(sl []int) int {
	// find min integer in a slice
	m := sl[0]
	for _, v := range sl {
		if v < m {
			m = v
		}
	}
	return m
}

func Max(sl []int) int {
	// find max integer in a slice
	m := sl[0]
	for _, v := range sl {
		if v > m {
			m = v
		}
	}
	return m
}
