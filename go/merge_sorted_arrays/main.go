package main

import "fmt"

func merge(nums1 []int, n int, nums2 []int, _ int) {
	var i, j, k int

	if n == len(nums1) {
		return
	}

	var buf = make([]int, 0, len(nums1))

	for {
		if i == len(nums1) {
			return
		}

		if nums1[i] == 11 {
			fmt.Print()
		}

		if i >= n {
			if j == len(nums2) {
				nums1[i], buf[k] = buf[k], nums1[i]
				k++
				i++
				continue
			}

			if k == len(buf) {
				nums1[i], nums2[j] = nums2[j], nums1[i]
				j++
				i++
				continue
			}

			if nums2[j] <= buf[k] {
				nums1[i], nums2[j] = nums2[j], nums1[i]
				j++
				i++
				continue
			}

			if buf[k] < nums2[j] {
				nums1[i], buf[k] = buf[k], nums1[i]
				k++
				i++
				continue
			}

		} else {
			if j == len(nums2) { // nums2 empty
				if k == len(buf) {
					continue
				}

				if buf[k] < nums1[i] {
					var tmp int
					nums1[i], tmp = buf[k], nums1[i]
					buf = append(buf, tmp)
					k++
					i++
					continue
				}

				i++
				continue
			}

			if k == len(buf) {
				if j == len(nums2) {
					continue
				}

				if nums2[j] < nums1[i] {
					var tmp int
					nums1[i], tmp = nums2[j], nums1[i]
					buf = append(buf, tmp)
					j++
				}
				i++
				continue
			}

			if nums2[j] < nums1[i] && nums2[j] <= buf[k] {
				var tmp int
				nums1[i], tmp = nums2[j], nums1[i]
				buf = append(buf, tmp)
				i++
				j++
				continue
			}

			if buf[k] < nums1[i] && buf[k] <= nums2[j] {
				var tmp int
				nums1[i], tmp = buf[k], nums1[i]
				buf = append(buf, tmp)
				k++
				i++
				continue
			}

			if nums1[i] <= buf[k] && nums1[i] <= nums2[j] {
				i++
			}
		}
	}
}
