package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	var root = new(ListNode)
	var out = root
	var prev = out

	for {
		bothNil := list1 == nil && list2 == nil
		bothNotNil := list1 != nil && list2 != nil
		onlyList1NotNil := list1 != nil && list2 == nil
		onlyList2NotNil := list2 != nil && list1 == nil

		if bothNil {
			out = nil
			break
		} else if onlyList1NotNil || bothNotNil && list1.Val <= list2.Val {
			*out = ListNode{
				Val: list1.Val,
			}
			list1 = list1.Next
		} else if onlyList2NotNil || bothNotNil && list2.Val <= list1.Val {
			*out = ListNode{
				Val: list2.Val,
			}
			list2 = list2.Next
		} else {
			panic("undefined state")
		}

		out.Next = new(ListNode)
		prev = out
		out = out.Next
	}

	prev.Next = nil

	return root
}
