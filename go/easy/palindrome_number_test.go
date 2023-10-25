package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome(121))
	assert.True(t, IsPalindrome(0))
	assert.True(t, IsPalindrome(1))
	assert.False(t, IsPalindrome(-121))
	assert.False(t, IsPalindrome(10))
}
