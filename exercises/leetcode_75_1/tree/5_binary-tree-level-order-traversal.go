/*
	Problem: 124. Binary Tree Maximum Path Sum
	Reference: https://leetcode.com/problems/binary-tree-level-order-traversal/
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

type Queue []*Node

func NewQueue() Queue {
	return Queue{}
}

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) Empty() bool {
	return q.Len() == 0
}

func (q *Queue) Push(nodes ...*Node) {
	for _, n := range nodes {
		if n != nil {
			*q = append(*q, n)
		}
	}
}

func (q *Queue) Pop() *Node {
	if q.Empty() {
		return nil
	}
	node := (*q)[0]
	*q = (*q)[1:]
	return node
}

type Node struct {
	Level     int
	reference *TreeNode
}

func (n *Node) Val() int {
	return n.reference.Val
}

func (n *Node) Left() *Node {
	if n.reference.Left == nil {
		return nil
	}
	return &Node{
		Level:     n.Level + 1,
		reference: n.reference.Left,
	}
}

func (n *Node) Right() *Node {
	if n.reference.Right == nil {
		return nil
	}
	return &Node{
		Level:     n.Level + 1,
		reference: n.reference.Right,
	}
}

func dfs(queue Queue, ans [][]int) [][]int {
	if queue.Empty() {
		return ans
	}
	node := queue.Pop()
	if len(ans) == node.Level {
		ans = append(ans, []int{node.Val()})
	} else {
		ans[node.Level] = append(ans[node.Level], node.Val())
	}
	queue.Push(node.Left(), node.Right())
	return dfs(queue, ans)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := NewQueue()
	queue.Push(&Node{Level: 0, reference: root})
	return dfs(queue, [][]int{})
}