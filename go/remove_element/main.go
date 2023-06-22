// https://leetcode.com/problems/remove-element/description/

package main

func removeElement(nums []int, val int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[r] == val {
			r--
			continue
		}

		if nums[l] == val {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
			continue
		}

		if nums[l] != val {
			l++
		}
	}

	return l
}
