/*
	Problem: 104. Maximum Depth of Binary Tree
	Reference: https://leetcode.com/problems/maximum-depth-of-binary-tree/
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

func calculateDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(calculateDepth(root.Left), calculateDepth(root.Right))
}

func maxDepth(root *TreeNode) int {
	return calculateDepth(root)
}