package easy

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	k := RemoveElement(nums, 3)
	assert.Equal(t, k, 2)
	sort.Ints(nums)
	assert.Equal(t, []int{2, 2}, nums[0:k])

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	k = RemoveElement(nums, 2)
	assert.Equal(t, k, 5)
	sort.Ints(nums[0:k])
	assert.Equal(t, []int{0, 0, 1, 3, 4}, nums[0:k])
}
