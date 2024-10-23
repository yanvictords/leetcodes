/*
	Problem: 572. Subtree of Another Tree
	Reference: https://leetcode.com/problems/subtree-of-another-tree
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

func findNode(node *TreeNode, subRoot *TreeNode, queue *[]*TreeNode) {
	if node == nil {
		return
	}
	if node.Val == subRoot.Val {
		*queue = append(*queue, node)
	}
	findNode(node.Left, subRoot, queue)
	findNode(node.Right, subRoot, queue)
}

func isEqual(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot != nil {
		return false
	}
	if root != nil && subRoot == nil {
		return false
	}
	if root == nil && subRoot == nil {
		return true
	}
	if root.Val != subRoot.Val {
		return false
	}
	return subtree(root.Left, subRoot.Left) &&
		subtree(root.Right, subRoot.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	roots := []*TreeNode{}
	findNode(root, subRoot, &roots)
	for _, r := range roots {
		if isEqual(r, subRoot) {
			return true
		}
	}
	return false
}