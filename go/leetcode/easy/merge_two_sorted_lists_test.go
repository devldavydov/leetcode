package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoSortedLists(t *testing.T) {
	assert.Equal(t, []int{1, 1, 2, 3, 4, 4}, sliceFromList(MergeTwoSortedLists(listFromSlice([]int{1, 2, 4}), listFromSlice([]int{1, 3, 4}))))
	assert.Equal(t, []int{}, sliceFromList(MergeTwoSortedLists(listFromSlice([]int{}), listFromSlice([]int{}))))
	assert.Equal(t, []int{0}, sliceFromList(MergeTwoSortedLists(listFromSlice([]int{}), listFromSlice([]int{0}))))
}

func listFromSlice(sl []int) *ListNode {
	var first, prev *ListNode
	for _, i := range sl {
		node := &ListNode{Val: i, Next: nil}
		if prev != nil {
			prev.Next = node
		}
		prev = node

		if first == nil {
			first = prev
		}
	}
	return first
}

func sliceFromList(p *ListNode) []int {
	sl := make([]int, 0)
	for p != nil {
		sl = append(sl, p.Val)
		p = p.Next
	}
	return sl
}
