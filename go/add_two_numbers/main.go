package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := intToList(123)
	l2 := intToList(321)

	l3 := addTwoNumbers(l1, l2)
	fmt.Println(listToInt(l3))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var root = new(ListNode)
	var node = root
	var overTen = false

	for {
		if l1 == nil && l2 == nil {
			if overTen {
				node.Next = &ListNode{Val: 1}
				break
			}

			break
		}

		v := 0
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}
		if overTen {
			v += 1
			overTen = false
		}
		if v > 9 {
			v %= 10
			overTen = true
		}

		node.Next = &ListNode{Val: v}
		node = node.Next
	}

	return root.Next
}

func listToInt(l *ListNode) int {
	var i = 1
	var r = 0
	for {
		if l == nil {
			break
		}
		r = r + l.Val*i
		i = i * 10
		l = l.Next
	}
	return r
}

func intToList(i int) *ListNode {
	if i == 0 {
		return new(ListNode)
	}

	var root = new(ListNode)
	var l = root
	for {
		if i == 0 {
			break
		}
		l2 := new(ListNode)
		l2.Val = i % 10
		l.Next = l2
		l = l2

		i = i / 10
	}
	return root.Next
}
