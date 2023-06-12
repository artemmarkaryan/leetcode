package main

func containsDuplicate(nums []int) bool {
	var m = make(map[int]int, len(nums))
	for i, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = i
	}
	return false
}
