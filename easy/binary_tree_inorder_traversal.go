package easy

import "github.com/devldavydov/leetcode/types"

func BinaryTreeInorderTraversal(root *types.TreeNode) []int {
	res := make([]int, 0)
	traverse(root, &res)
	return res
}

func traverse(node *types.TreeNode, list *[]int) {
	if node == nil {
		return
	}

	traverse(node.Left, list)
	*list = append(*list, node.Val)
	traverse(node.Right, list)
}
