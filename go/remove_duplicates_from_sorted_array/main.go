package main

import "fmt"

func removeDuplicates(nums []int) int {
	i := 0
	for {
		if i == len(nums)-1 {
			break
		}

		if nums[i+1] > nums[i] {
			i++
			continue
		}

		// nums[i+1] == nums[i]
		for j := i + i; j < len(nums); j++ {
			if nums[j] > nums[i] {
				nums[i+1], nums[j] = nums[j], nums[i+1]
				break
			}
		}

		i++
	}

	fmt.Println(nums)
	fmt.Println(i + 1)

	return i + 1
}
