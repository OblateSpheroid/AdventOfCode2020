// Some generic helper functions be used by different days
package helpers

import (
	"sort"
)

// Return a copy of an integer slice
func CopySlice(sl []int) []int {
	tmp := make([]int, len(sl))
	for i, v := range sl {
		tmp[i] = v
	}
	return tmp
}

// Return a copy of a slice that is sorted
func SortSlice(sl []int) []int {
	sorted := sort.IntSlice(CopySlice(sl))
	sort.Sort(sorted)
	return sorted
}

// Test if an integer is in a sorted integer slice
func IsInSorted(i int, sl []int) bool {
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

// Test if an integer is in a slice
func IsIn(i int, sl []int) bool {
	for _, v := range sl {
		if i == v {
			return true
		}
	}
	return false
}

// Test if a string is in a slice
func IsInString(s string, sl []string) bool {
	for _, v := range sl {
		if s == v {
			return true
		}
	}
	return false
}

// Append an integer string "in place", without returning a new slice
func AppendInPlace(sl *[]int, i ...int) {
	for _, j := range i {
		*sl = append(*sl, j)
	}
}

// Make a sequence of numbers based on a min and max
func MakeSeq(min, max int) []int {
	sl := []int{}
	for i := min; i <= max; i++ {
		AppendInPlace(&sl, i)
	}
	return sl
}

// Return a copy of a slice without a specific integer included as an element
func Drop(sl []int, i int) []int {
	new := []int{}
	for _, v := range sl {
		if v != i {
			new = append(new, v)
		}
	}
	return new
}

// Return a copy of a slice without a specific integer included as an element
func DropString(sl []string, s string) []string {
	new := []string{}
	for _, v := range sl {
		if v != s {
			new = append(new, v)
		}
	}
	return new
}

// Test if string is a key in a map
func IsInMap(key string, m map[string]int) bool {
	for k := range m {
		if k == key {
			return true
		}
	}
	return false
}

// Convert map keys to a slice of unique strings
func MapToSlice(m map[string]int) []string {
	s := []string{}
	for k := range m {
		s = append(s, k)
	}
	return s
}

// Some basic math functions

// Sum all integers in a slice
func Sum(sl []int) int {
	tot := 0
	for _, v := range sl {
		tot += v
	}
	return tot
}

// Find min integer in a slice
func Min(sl []int) int {
	m := sl[0]
	for _, v := range sl {
		if v < m {
			m = v
		}
	}
	return m
}

// Find max integer in a slice
func Max(sl []int) int {
	m := sl[0]
	for _, v := range sl {
		if v > m {
			m = v
		}
	}
	return m
}
