/*
	Problem: 100. Same Tree
	Reference: https://leetcode.com/problems/same-tree/
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

func isSameNode(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameNode(p.Left, q.Left) && isSameNode(p.Right, q.Right)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameNode(p, q)
}