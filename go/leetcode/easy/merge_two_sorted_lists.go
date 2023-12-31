package easy

/*
https://leetcode.com/problems/merge-two-sorted-lists/

You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.
*/

import "github.com/devldavydov/leetcode/types"

func MergeTwoSortedLists(list1 *types.ListNode, list2 *types.ListNode) *types.ListNode {
	var prev, first *types.ListNode

loop:
	for {
		switch {
		case list1 != nil && list2 != nil:
			if list1.Val <= list2.Val {
				prev = addNode(prev, list1.Val)
				list1 = list1.Next
			} else {
				prev = addNode(prev, list2.Val)
				list2 = list2.Next
			}
		case list1 != nil && list2 == nil:
			prev = addNode(prev, list1.Val)
			list1 = list1.Next
		case list1 == nil && list2 != nil:
			prev = addNode(prev, list2.Val)
			list2 = list2.Next
		case list1 == nil && list2 == nil:
			break loop
		}

		if first == nil {
			first = prev
		}
	}

	return first
}

func addNode(prev *types.ListNode, val int) *types.ListNode {
	newNode := &types.ListNode{Val: val, Next: nil}
	if prev != nil {
		prev.Next = newNode
	}
	return newNode
}
