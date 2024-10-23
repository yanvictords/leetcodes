/*
	Problem: 124. Binary Tree Maximum Path Sum
	Reference: https://leetcode.com/problems/binary-tree-maximum-path-sum/
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

func nonNegative(num int) int {
	return max(num, 0)
}

func dfs(node *TreeNode, ans *int) int {
	if node == nil {
		return 0
	}
	curr := node.Val
	leftAcc := dfs(node.Left, ans)
	rightAcc := dfs(node.Right, ans)
	horizontalSum := curr + nonNegative(leftAcc) + nonNegative(rightAcc)
	*ans = max(*ans, horizontalSum)
	return curr + max(nonNegative(leftAcc), nonNegative(rightAcc))
}

func maxPathSum(root *TreeNode) int {
	ans := -1001
	dfs(root, &ans)
	return ans
}