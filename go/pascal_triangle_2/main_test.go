package main

import (
	"reflect"
	"testing"
)

func Test_getRow(t *testing.T) {
	tests := []struct {
		i int
		w []int
	}{
		{i: 0, w: []int{1}},
		{i: 1, w: []int{1, 1}},
		{i: 2, w: []int{1, 2, 1}},
		{i: 3, w: []int{1, 3, 3, 1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := getRow(tt.i); !reflect.DeepEqual(got, tt.w) {
				t.Errorf("getRow(%v) = %v, want %v", tt.i, got, tt.w)
			}
		})
	}
}
