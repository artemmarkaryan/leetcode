package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{
			s: "ab_a",
			want: true,
		},
		{
			s:    "0P",
			want: false,
		},
		{
			s:    "race a car",
			want: false,
		},
		{
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
