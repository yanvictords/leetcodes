/*
	Problem: 297. Serialize and Deserialize Binary Tree
	Reference: https://leetcode.com/problems/serialize-and-deserialize-binary-tree
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

type Nodes []string

func NewNodes(serializable string) Nodes {
	splitted := strings.Split(serializable, ",")
	return splitted
}

func (n *Nodes) Len() int {
	return len(*n)
}

func (n *Nodes) Empty() bool {
	return n.Len() == 0
}

func (n *Nodes) Pop() *int {
	if n.Empty() {
		return nil
	}
	node := (*n)[0]
	*n = (*n)[1:]
	if node == "null" {
		return nil
	}
	num, err := strconv.Atoi(node)
	if err != nil {
		return nil
	}
	return &num
}

type Queue []*TreeNode

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) Empty() bool {
	return q.Len() == 0
}

func (q *Queue) Push(node *TreeNode) {
	*q = append(*q, node)

}

func (q *Queue) Pop() *TreeNode {
	if q.Empty() {
		return nil
	}
	next := (*q)[0]
	*q = (*q)[1:]
	return next
}

type Codec struct {
	queue Queue
}

func Constructor() Codec {
	return Codec{
		queue: Queue{},
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	queue := Queue{}
	queue.Push(root)
	result := this.toString(queue)
	return result[:len(result)-1]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	queue := Queue{}
	nodes := NewNodes(data)
	rootVal := nodes.Pop()
	if rootVal == nil {
		return nil
	}
	root := &TreeNode{Val: *rootVal}
	queue.Push(root)
	this.buildTree(queue, nodes)
	return root
}

func (this *Codec) buildTree(queue Queue, nodes Nodes) {
	if queue.Empty() {
		return
	}
	root := queue.Pop()
	if root == nil {
		return
	}

	if l := nodes.Pop(); l != nil {
		left := &TreeNode{Val: *l}
		root.Left = left
		queue.Push(left)
	}

	if r := nodes.Pop(); r != nil {
		right := &TreeNode{Val: *r}
		root.Right = right
		queue.Push(right)
	}

	this.buildTree(queue, nodes)
}

func (this *Codec) toString(queue Queue) string {
	if queue.Empty() {
		return ""
	}
	next := queue.Pop()
	if next == nil {
		return "null," + this.toString(queue)
	}
	queue.Push(next.Left)
	queue.Push(next.Right)
	return fmt.Sprintf("%d,", next.Val) + this.toString(queue)
}

/**
* Your Codec object will be instantiated and called as such:
* ser := Constructor();
* deser := Constructor();
* data := ser.serialize(root);
* ans := deser.deserialize(data);
 */