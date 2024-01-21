package easy

import (
	"testing"

	"github.com/devldavydov/leetcode/types"
	"github.com/stretchr/testify/assert"
)

func TestMergeTwoSortedLists(t *testing.T) {
	assert.Equal(t,
		[]int{1, 1, 2, 3, 4, 4},
		types.SliceFromList(
			MergeTwoSortedLists(
				types.ListFromSlice([]int{1, 2, 4}),
				types.ListFromSlice([]int{1, 3, 4}))))
	assert.Equal(t, []int{},
		types.SliceFromList(
			MergeTwoSortedLists(
				types.ListFromSlice([]int{}),
				types.ListFromSlice([]int{}))))
	assert.Equal(t,
		[]int{0},
		types.SliceFromList(
			MergeTwoSortedLists(
				types.ListFromSlice([]int{}),
				types.ListFromSlice([]int{0}))))
}
