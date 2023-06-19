package main

import (
	"fmt"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 int
		l2 int
	}
	tests := []struct {
		args args
		want int
	}{
		// {
		// 	args: args{
		// 		l1: 123,
		// 		l2: 321,
		// 	},
		// 	want: 444,
		// },
		{
			args: args{
				l1: 9,
				l2: 9,
			},
			want: 18,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := addTwoNumbers(
				intToList(tt.args.l1),
				intToList(tt.args.l2),
			); listToInt(got) != tt.want {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
