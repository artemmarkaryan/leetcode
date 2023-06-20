// https://leetcode.com/problems/longest-palindromic-substring/

package main

import (
	"fmt"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{in: "avabv", out: "ava"},
		{in: "abcdc", out: "cdc"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			if got := longestPalindrome(tt.in); got != tt.out {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.out)
			}
		})
	}
}
