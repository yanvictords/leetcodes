/*
	Problem: 23. Merge k Sorted Lists
	Reference: https://leetcode.com/problems/merge-k-sorted-lists/
    Complexity: O(N log n)
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type Stream struct {
	Lists []*ListNode
}

func (k *Stream) Merge() *ListNode {
	return k.flat().
		sort().
		join()
}

func (k *Stream) flat() *Stream {
	flat := []*ListNode{}
	for _, head := range k.Lists {
		aux := head
		for aux != nil {
			flat = append(flat, aux)
			next := aux.Next
			aux.Next = nil
			aux = next
		}
	}
	k.Lists = flat
	return k
}

func (k *Stream) sort() *Stream {
	k.Lists = mergeSort(k.Lists)
	return k
}

func (k *Stream) join() *ListNode {
	var head, next *ListNode
	for _, node := range k.Lists {
		if head == nil {
			head = node
			next = head
			continue
		}
		next.Next = node
        next = next.Next
	}
	return head
}

func mergeKLists(lists []*ListNode) *ListNode {
	stream := Stream{
		Lists: lists,
	}
	return stream.Merge()
}

func mergeSort(lists []*ListNode) []*ListNode {
	sz := len(lists)
	if sz < 2 {
		return lists
	}
	listA := mergeSort(lists[:sz/2])
	listB := mergeSort(lists[sz/2:])
	return merge(listA, listB)
}

func merge(listA []*ListNode, listB []*ListNode) []*ListNode {
	sorted := []*ListNode{}
	i, j := 0, 0
	for i < len(listA) && j < len(listB) {
		if listA[i].Val < listB[j].Val {
			sorted = append(sorted, listA[i])
			i++
			continue
		}
		sorted = append(sorted, listB[j])
		j++
		continue
	}

	for ; i < len(listA); i++ {
		sorted = append(sorted, listA[i])
	}
	for ; j < len(listB); j++ {
		sorted = append(sorted, listB[j])
	}

	return sorted
}