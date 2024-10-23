/*
	Problem: 23. Merge k Sorted Lists
	Reference: https://leetcode.com/problems/merge-k-sorted-lists/
    Complexity: O(N log K)
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type Heap struct {
    Tree []*ListNode
}

func NewHeap() *Heap {
    return &Heap {
        Tree: make([]*ListNode, 0),
    }
}

func (h *Heap) Len() int {
    return len(h.Tree)
}

func (h *Heap) Swap(i, j int) {
    h.Tree[i], h.Tree[j] = h.Tree[j], h.Tree[i]
}

func (h *Heap) Push(node *ListNode) {
    h.Tree = append(h.Tree, node)
    h.siftUp(h.Len()-1)
}

func (h *Heap) Pop() *ListNode {
    if h.Len() == 0 {
        return nil
    }
    node := h.Tree[0]
    h.Tree[0] = h.Tree[h.Len()-1]
    h.Tree = h.Tree[:h.Len()-1]
    h.siftDown(0)
    return node
}

func (h *Heap) siftUp(idx int) {
    for idx > 0 {
        parent := (idx - 1) / 2
        if h.Tree[idx].Val > h.Tree[parent].Val {
            break
        }
        h.Swap(idx, parent)
        idx = parent
    }
}

func (h *Heap) siftDown(idx int) {
    for idx < h.Len() {
        left := (idx * 2) + 1
        right := (idx * 2) + 2
        children := left
        if right < h.Len() && h.Tree[left].Val > h.Tree[right].Val {
            children = right
        }
        if children >= h.Len() || h.Tree[idx].Val < h.Tree[children].Val {
            break
        }
        h.Swap(idx, children)
        idx = children
    }
}

func mergeKLists(lists []*ListNode) *ListNode {
    heap := NewHeap()
    for _, l := range lists {
        curr := l
        for curr != nil {
            aux := curr
            curr = curr.Next
            aux.Next = nil
            heap.Push(aux)
        }
    }

    var head, ans *ListNode
    sz := heap.Len()
    for i := 0; i < sz; i++ {
        if ans == nil {
            head = heap.Pop()
            ans = head
            continue
        }
        ans.Next = heap.Pop()
        ans = ans.Next
    }

    return head
}