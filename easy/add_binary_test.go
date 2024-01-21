package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBinary(t *testing.T) {
	assert.Equal(t, "100", AddBinary("11", "1"))
	assert.Equal(t, "10101", AddBinary("1010", "1011"))
	assert.Equal(t, "1000", AddBinary("1", "111"))
}
