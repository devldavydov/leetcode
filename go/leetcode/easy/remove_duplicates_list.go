package easy

import "github.com/devldavydov/leetcode/types"

/*
https://leetcode.com/problems/remove-duplicates-from-sorted-list/

Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.
*/

func RemoveDuplicatesFromSortedList(head *types.ListNode) *types.ListNode {
	p, l := head, head
	for p != nil {
		if l.Val != p.Val {
			l = p
			p = p.Next
		} else {
			t := p
			p = p.Next
			l.Next = t.Next
		}
	}
	return head
}
