package main

// Some basic math functions

func sum(sl []int) int {
	// sum all integers in a slice
	tot := 0
	for _, v := range sl {
		tot += v
	}
	return tot
}

func min(sl []int) int {
	// find min integer in a slice
	m := sl[0]
	for _, v := range sl {
		if v < m {
			m = v
		}
	}
	return m
}

func max(sl []int) int {
	// find max integer in a slice
	m := sl[0]
	for _, v := range sl {
		if v > m {
			m = v
		}
	}
	return m
}
