/*
	Problem: 235. Lowest Common Ancestor of a Binary Search Tree
	Reference: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
    Complexity: O(N)
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

 func dfs(node *TreeNode, p, q int) *TreeNode {
    if node == nil {
        return nil
    }
    if node.Val == p || node.Val == q {
        return node
    }
    left := dfs(node.Left, p, q)
    right := dfs(node.Right, p, q)
    if left != nil && right != nil {
        return node
    }
    if left != nil {
        return left
    }
    return right
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return dfs(root, p.Val, q.Val)
}