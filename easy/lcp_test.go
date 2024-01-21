package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	assert.Equal(t, "fl", LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	assert.Equal(t, "", LongestCommonPrefix([]string{"dog", "racecar", "car"}))
}
