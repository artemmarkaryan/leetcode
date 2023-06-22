// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

package main

func removeDuplicates(n []int) int {
	i := 1
	prev := n[0]
	for j := 1; j < len(n); j++ {
		if n[j] > prev {
			n[i] = n[j]
			i++
		}
		prev = n[j]
	}

	return i
}
