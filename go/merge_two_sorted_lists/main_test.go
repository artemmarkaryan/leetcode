package main

import (
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				list1: &ListNode{
					Val:  1,
					Next: nil,
				},
				list2: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.list1, tt.args.list2)
			compare(t, got, tt.want)
		})
	}
}

func compare(t *testing.T, a, b *ListNode) {
	for {
		if a.Val != b.Val {
			t.Errorf("not equal: a.Val = %v, b.Val = %v", a.Val, b.Val)
			t.FailNow()
		}

		if (a.Next == nil) && (b.Next == nil) {
			return
		}

		if (a.Next != nil) && (b.Next == nil) {
			t.Error("a.Next != nil && b.Next == nil")
			t.FailNow()
		}

		if (a.Next == nil) && (b.Next != nil) {
			t.Error("a.Next == nil && b.Next != nil")
			t.FailNow()
		}

		if (a.Next != nil) && (b.Next != nil) {
			a = a.Next
			b = b.Next
		}
	}
}
