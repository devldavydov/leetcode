package easy

import (
	"testing"

	"github.com/devldavydov/leetcode/types"
	"github.com/stretchr/testify/assert"
)

func TestIsSameTree(t *testing.T) {
	assert.True(t, IsSameTree(
		&types.TreeNode{
			Val: 1,
			Left: &types.TreeNode{
				Val: 2,
			},
			Right: &types.TreeNode{
				Val: 3,
			},
		},
		&types.TreeNode{
			Val: 1,
			Left: &types.TreeNode{
				Val: 2,
			},
			Right: &types.TreeNode{
				Val: 3,
			},
		},
	))
	assert.False(t, IsSameTree(
		&types.TreeNode{
			Val: 1,
			Left: &types.TreeNode{
				Val: 2,
			},
		},
		&types.TreeNode{
			Val: 1,
			Right: &types.TreeNode{
				Val: 2,
			},
		},
	))
	assert.False(t, IsSameTree(
		&types.TreeNode{
			Val: 1,
			Left: &types.TreeNode{
				Val: 2,
			},
			Right: &types.TreeNode{
				Val: 3,
			},
		},
		&types.TreeNode{
			Val: 1,
			Left: &types.TreeNode{
				Val: 3,
			},
			Right: &types.TreeNode{
				Val: 2,
			},
		},
	))
}
