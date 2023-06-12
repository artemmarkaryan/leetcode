package main

import (
	"sort"
)

// O(nlogn)
func hIndex(citations []int) int {
	sort.Slice(citations, func(i int, j int) bool { return citations[i] > citations[j] })
	for i := 0; i < len(citations); i++ {
		if i+1 > citations[i] {
			return i
		}
	}

	return len(citations)
}

// O(n)
func hIndex2(citations []int) int {
	c := [5000]int{}
	for _, v := range citations {
		if v > len(citations) {
			c[len(citations)] += 1
		} else {
			c[v] += 1
		}
	}

	var r = make([]int, 0, len(citations))
	for i, v := range c {
		for j := 0; j < v; j++ {
			r = append(r, i)
		}
	}

	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-i-1] = r[len(r)-i-1], r[i]
	}

	for i := 0; i < len(r); i++ {
		if i+1 > r[i] {
			return i
		}
	}

	return 0
}
