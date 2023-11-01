package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidParentheses(t *testing.T) {
	assert.True(t, IsValidParentheses("()"))
	assert.True(t, IsValidParentheses("()[]{}"))
	assert.True(t, IsValidParentheses("(({{[[()]]}}))"))
	assert.False(t, IsValidParentheses("(]"))
	assert.False(t, IsValidParentheses("("))
	assert.False(t, IsValidParentheses(")"))
}
