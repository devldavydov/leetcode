package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInteger(t *testing.T) {
	assert.Equal(t, 3, RomanToInt("III"))
	assert.Equal(t, 58, RomanToInt("LVIII"))
	assert.Equal(t, 1994, RomanToInt("MCMXCIV"))
}
