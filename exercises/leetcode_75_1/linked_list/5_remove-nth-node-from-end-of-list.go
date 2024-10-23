/*
	Problem: 19. Remove Nth Node From End of List
	Reference: https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
    Complexity: O(n)
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	mapper := make(map[int]*ListNode)
	aux := head
	total := 0
	for i := 0; aux != nil; i++ {
		total++
		mapper[i] = aux
		aux = aux.Next
	}
	sz := len(mapper)
	n = sz - n
	if n == 0 {
		return mapper[1]
	}
	if n == sz-1 {
		mapper[n-1].Next = nil
		return head
	}
	mapper[n-1].Next = mapper[n+1]
	return head
}