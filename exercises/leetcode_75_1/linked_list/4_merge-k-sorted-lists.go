/*
	Problem: 23. Merge k Sorted Lists
	Reference: https://leetcode.com/problems/merge-k-sorted-lists/description/
    Complexity: O(Kn)
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func merge(left []*ListNode, right []*ListNode) []*ListNode {
	ans := []*ListNode{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i].Val < right[j].Val {
			ans = append(ans, left[i])
			i++
			continue
		}
		ans = append(ans, right[j])
		j++
	}
	for ; i < len(left); i++ {
		ans = append(ans, left[i])
	}
	for ; j < len(right); j++ {
		ans = append(ans, right[j])
	}
	return ans
}

func mergeSort(lists []*ListNode) []*ListNode {
	if len(lists) < 2 {
		return lists
	}
	sz := len(lists) / 2
	l := mergeSort(lists[:sz])
	r := mergeSort(lists[sz:])
	return merge(l, r)
}

func flatLists(lists []*ListNode) []*ListNode {
	flat := []*ListNode{}
	for _, root := range lists {
		for root != nil {
			flat = append(flat, &ListNode{Val: root.Val})
			root = root.Next
		}
	}

	return flat
}

func linkNodes(nodes []*ListNode) *ListNode {
	var root, aux *ListNode
	for _, node := range nodes {
		if aux == nil {
			aux = node
			root = aux
			continue
		}
		aux.Next = node
		aux = aux.Next
	}
	return root
}

func mergeKLists(lists []*ListNode) *ListNode {
	flatted := flatLists(lists)
	sorted := mergeSort(flatted)
	return linkNodes(sorted)
}