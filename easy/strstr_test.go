package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	assert.Equal(t, 0, StrStr("sadbutsad", "sad"))
	assert.Equal(t, 6, StrStr("sedbutsad", "sad"))
	assert.Equal(t, -1, StrStr("leetcode", "leeto"))
}
