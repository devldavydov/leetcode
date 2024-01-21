package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	nums := []int{1, 1, 2}
	k := RemoveDuplicatesFromSortedArray(nums)
	assert.Equal(t, 2, k)
	assert.Equal(t, []int{1, 2}, nums[0:k])

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k = RemoveDuplicatesFromSortedArray(nums)
	assert.Equal(t, 5, k)
	assert.Equal(t, []int{0, 1, 2, 3, 4}, nums[0:k])
}
