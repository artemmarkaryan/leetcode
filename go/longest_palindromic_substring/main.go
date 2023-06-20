// https://leetcode.com/problems/longest-palindromic-substring/

package main

func longestPalindrome(s string) string {
	maxPal := [2]int{}
	for i := 0; i < len(s); i++ {
		var l, r int
		for l, r = i, i; l >= 0 && r < len(s); func() { l--; r++ }() {
			if s[l] == s[r] {
				if r-l > maxPal[1]-maxPal[0] {
					maxPal[0], maxPal[1] = l, r
				}
			} else {
				break
			}
		}
		for l, r = i, i+1; l >= 0 && r < len(s); func() { l--; r++ }() {
			if s[l] == s[r] {
				if r-l > maxPal[1]-maxPal[0] {
					maxPal[0], maxPal[1] = l, r
				}
			} else {
				break
			}
		}
	}
	return s[maxPal[0] : maxPal[1]+1]
}
