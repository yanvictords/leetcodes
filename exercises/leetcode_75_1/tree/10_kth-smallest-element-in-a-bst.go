/*
	Problem: 230. Kth Smallest Element in a BST
	Reference: https://leetcode.com/problems/kth-smallest-element-in-a-bst/
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

 func buildArr(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    left := buildArr(root.Left)
    mid := []int{root.Val}
    right := buildArr(root.Right)
    return append(left, append(mid, right...)...)
}

func kthSmallest(root *TreeNode, k int) int {
    return buildArr(root)[k-1]
}