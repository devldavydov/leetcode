package types

func SliceFromList(p *ListNode) []int {
	sl := make([]int, 0)
	for p != nil {
		sl = append(sl, p.Val)
		p = p.Next
	}
	return sl
}

func ListFromSlice(sl []int) *ListNode {
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
