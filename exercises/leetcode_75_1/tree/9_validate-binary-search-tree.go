/*
	Problem: 98. Validate Binary Search Tree
	Reference: https://leetcode.com/problems/validate-binary-search-tree/
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

func isSorted(arr []int) bool {
	if len(arr) == 0 {
		return true
	}
	left := arr[0]
	for i := 1; i < len(arr); i++ {
		right := arr[i]
		if left >= right {
			return false
		}
		left = right
	}
	return true
}

func buildArr(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := buildArr(root.Left)
	mid := []int{root.Val}
	right := buildArr(root.Right)
	return append(left, append(mid, right...)...)
}

func isValidBST(root *TreeNode) bool {
	return isSorted(buildArr(root))
}