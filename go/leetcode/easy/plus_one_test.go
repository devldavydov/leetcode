package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	assert.Equal(t, []int{1}, PlusOne([]int{0}))
	assert.Equal(t, []int{1, 2, 4}, PlusOne([]int{1, 2, 3}))
	assert.Equal(t, []int{4, 3, 2, 2}, PlusOne([]int{4, 3, 2, 1}))
	assert.Equal(t, []int{1, 0}, PlusOne([]int{9}))
	assert.Equal(t, []int{1, 0, 0, 0, 0}, PlusOne([]int{9, 9, 9, 9}))
}
