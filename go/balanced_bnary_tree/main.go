package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var md = -1
var balanced = true

func w(n *TreeNode, i int) {
	if !balanced {
		return
	}
	if n.Left == nil && n.Right == nil {
		if md == -1 {
			md = i
		} else {
			if md > i {
				if md-i > 1 {
					balanced = false
					return
				} else {
					md = i
				}
			}
			if md < i && i-md > 1 {
				balanced = false
				return
			}
		}
		return
	}
	if n.Left != nil {
		w(n.Left, i+1)
	}
	if n.Right != nil {
		w(n.Right, i+1)
	}
}

func isBalanced(root *TreeNode) bool {
	if root != nil {
		w(root, 0)
		return balanced
	}
	return true
}
