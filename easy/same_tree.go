package easy

import "github.com/devldavydov/leetcode/types"

/*
https://leetcode.com/problems/same-tree/

Given the roots of two binary trees p and q, write a function to check if they are the same or not.
Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
*/

func IsSameTree(p *types.TreeNode, q *types.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if !(p != nil && q != nil && p.Val == q.Val) {
		return false
	}

	if !IsSameTree(p.Left, q.Left) {
		return false
	}

	return IsSameTree(p.Right, q.Right)
}
