package easy

import (
	"testing"

	"github.com/devldavydov/leetcode/types"
	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicatesFromSortedList(t *testing.T) {
	assert.Equal(t,
		[]int{1, 2},
		types.SliceFromList(
			RemoveDuplicatesFromSortedList(
				types.ListFromSlice([]int{1, 1, 2}))))
	assert.Equal(t,
		[]int{1, 2, 3},
		types.SliceFromList(
			RemoveDuplicatesFromSortedList(
				types.ListFromSlice([]int{1, 1, 2, 3, 3}))))
	assert.Equal(t,
		[]int{},
		types.SliceFromList(
			RemoveDuplicatesFromSortedList(
				types.ListFromSlice([]int{}))))
}
