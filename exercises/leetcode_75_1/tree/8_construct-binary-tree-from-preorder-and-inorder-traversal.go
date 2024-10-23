/*
	Problem: 105. Construct Binary Tree from Preorder and Inorder Traversal
	Reference: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
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
func build(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	idx := 0
	for i, val := range inorder {
		if val == rootVal {
			idx = i
			break
		}
	}
	lpreorder, rpreorder := preorder[1:idx+1], preorder[idx+1:]
	linorder, rinorder := inorder[:idx], inorder[idx+1:]

	return &TreeNode{
		Val:   rootVal,
		Left:  build(lpreorder, linorder),
		Right: build(rpreorder, rinorder),
	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, inorder)
}