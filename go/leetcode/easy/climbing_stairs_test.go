package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbingStairs(t *testing.T) {
	assert.Equal(t, 2, ClimbingStairs(2))
	assert.Equal(t, 3, ClimbingStairs(3))
}
