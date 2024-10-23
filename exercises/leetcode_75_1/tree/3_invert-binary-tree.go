/*
	Problem: 226. Invert Binary Tree
	Reference: https://leetcode.com/problems/invert-binary-tree/
    Complexity: O(N)
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertNode(node *TreeNode) {
	if node == nil {
		return
	}
	node.Left, node.Right = node.Right, node.Left
	invertNode(node.Left)
	invertNode(node.Right)
}

func invertTree(root *TreeNode) *TreeNode {
	invertNode(root)
	return root
}