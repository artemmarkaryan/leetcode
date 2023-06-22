// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string

package main

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle); i++ {
		found := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				found = false
				break
			}
		}

		if found {
			return i
		}
	}

	return -1
}
