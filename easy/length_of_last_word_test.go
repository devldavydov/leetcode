package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	assert.Equal(t, 5, LengthOfLastWord("Hello World"))
	assert.Equal(t, 4, LengthOfLastWord("   fly me   to   the moon  "))
	assert.Equal(t, 6, LengthOfLastWord("luffy is still joyboy"))
}
