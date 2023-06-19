package main

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			in:  "abc",
			out: 3,
		},
		{
			in:  "abca",
			out: 3,
		},
		{
			in:  "abad",
			out: 3,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.in); got != tt.out {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.out)
			}
		})
	}
}
