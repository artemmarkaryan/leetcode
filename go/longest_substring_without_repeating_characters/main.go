package main

import "strings"

func lengthOfLongestSubstring(s string) int {
	var maxSubstringLength = 0
	var substring = new(strings.Builder)
	var substringCache = make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		if ii, ok := substringCache[s[i]]; ok {
			for k := range substringCache {
				delete(substringCache, k)
			}
			if substring.Len() > maxSubstringLength {
				maxSubstringLength = substring.Len()
			}
			substring = new(strings.Builder)
			i = ii + 1
		}

		substringCache[s[i]] = i
		substring.WriteByte(s[i])
	}

	if substring.Len() > maxSubstringLength {
		maxSubstringLength = substring.Len()
	}

	return maxSubstringLength
}
