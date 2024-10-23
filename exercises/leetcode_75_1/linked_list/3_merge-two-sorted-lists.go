/*
	Problem: 21. Merge Two Sorted Lists
	Reference: https://leetcode.com/problems/merge-two-sorted-lists/
    Complexity: O(n)
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func cloneNode(node *ListNode) *ListNode {
	return &ListNode{
		Val: node.Val,
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var root, aux *ListNode

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			if aux == nil {
				aux = cloneNode(list1)
				root = aux
			} else {
				aux.Next = cloneNode(list1)
				aux = aux.Next
			}
			list1 = list1.Next
			continue
		}
		if aux == nil {
			aux = cloneNode(list2)
			root = aux
		} else {
			aux.Next = cloneNode(list2)
			aux = aux.Next
		}
		list2 = list2.Next
	}

	for list1 != nil {
		if aux == nil {
			aux = cloneNode(list1)
			root = aux
			list1 = list1.Next
			continue
		}
		aux.Next = cloneNode(list1)
		aux = aux.Next
		list1 = list1.Next
	}

	for list2 != nil {
		if aux == nil {
			aux = cloneNode(list2)
			root = aux
			list2 = list2.Next
			continue
		}
		aux.Next = cloneNode(list2)
		aux = aux.Next
		list2 = list2.Next
	}

	return root
}