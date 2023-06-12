package main

import (
	"fmt"
	"testing"
)

func Test_hIndex(t *testing.T) {
	tests := []struct {
		citations []int
		want      int
	}{
		{
			citations: []int{6, 5, 3, 1},
			want:      3,
		},
		{
			citations: []int{3, 1, 1},
			want:      1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := hIndex(tt.citations); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
			if got := hIndex2(tt.citations); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
