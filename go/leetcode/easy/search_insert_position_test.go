package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsertPosition(t *testing.T) {
	assert.Equal(t, 2, SearchInsertPosition([]int{1, 3, 5, 6}, 5))
	assert.Equal(t, 1, SearchInsertPosition([]int{1, 3, 5, 6}, 2))
	assert.Equal(t, 4, SearchInsertPosition([]int{1, 3, 5, 6}, 7))
	assert.Equal(t, 0, SearchInsertPosition([]int{1, 3, 5, 6}, 0))
}
