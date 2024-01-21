package easy

import (
	"testing"

	"github.com/devldavydov/leetcode/types"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeInorderTraversal(t *testing.T) {
	assert.Equal(t, []int{1, 3, 2}, BinaryTreeInorderTraversal(
		&types.TreeNode{
			Val: 1,
			Right: &types.TreeNode{
				Val: 2,
				Left: &types.TreeNode{
					Val: 3,
				},
			},
		}))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, BinaryTreeInorderTraversal(
		&types.TreeNode{
			Val: 6,
			Left: &types.TreeNode{
				Val: 2,
				Left: &types.TreeNode{
					Val: 1,
				},
				Right: &types.TreeNode{
					Val: 4,
					Left: &types.TreeNode{
						Val: 3,
					},
					Right: &types.TreeNode{
						Val: 5,
					},
				},
			},
			Right: &types.TreeNode{
				Val: 7,
				Right: &types.TreeNode{
					Val: 9,
					Left: &types.TreeNode{
						Val: 8,
					},
				},
			},
		},
	))
	assert.Equal(t, []int{}, BinaryTreeInorderTraversal(nil))
	assert.Equal(t, []int{1}, BinaryTreeInorderTraversal(&types.TreeNode{Val: 1, Left: nil, Right: nil}))
}
