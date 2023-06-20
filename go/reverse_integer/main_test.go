// https://leetcode.com/problems/reverse-integer/

package main

import (
	"fmt"
	"testing"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		in, out int
	}{
		{1463847412, 2147483641},
		{1534236469, 0},
		{123, 321},
		{-123, -321},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			if gotR := reverse(tt.in); gotR != tt.out {
				t.Errorf("reverse() = %v, want %v", gotR, tt.out)
			}
		})
	}
}
