package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqrtx(t *testing.T) {
	assert.Equal(t, 0, Sqrtx(0))
	assert.Equal(t, 1, Sqrtx(1))
	assert.Equal(t, 1, Sqrtx(2))
	assert.Equal(t, 1, Sqrtx(3))
	assert.Equal(t, 2, Sqrtx(4))
	assert.Equal(t, 2, Sqrtx(8))
	assert.Equal(t, 3, Sqrtx(9))
	assert.Equal(t, 4, Sqrtx(16))
	assert.Equal(t, 4, Sqrtx(18))
	assert.Equal(t, 25, Sqrtx(625))
}
