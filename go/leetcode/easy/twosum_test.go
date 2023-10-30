package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	res := TwoSum([]int{2, 7, 11, 15}, 9)
	assert.True(t, containsInSlice(res, 0) && containsInSlice(res, 1))

	res = TwoSum([]int{3, 2, 4}, 6)
	assert.True(t, containsInSlice(res, 1) && containsInSlice(res, 2))

	res = TwoSum([]int{3, 3}, 6)
	assert.True(t, containsInSlice(res, 0) && containsInSlice(res, 1))
}

func containsInSlice(s []int, i int) bool {
	for _, item := range s {
		if item == i {
			return true
		}
	}
	return false
}
