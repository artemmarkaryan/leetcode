package main

func containsNearbyDuplicate(nums []int, k int) bool {
	var m = make(map[int]int, len(nums))

	for i, v := range nums {
		if j, ok := m[v]; ok && diff(i, j) <= k {
			return true
		}
		m[v] = i
	}

	return false
}

func diff(i, j int) int {
	v := i - j
	if v < 0 {
		return -v
	}
	return v
}
