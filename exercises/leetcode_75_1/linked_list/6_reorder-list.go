/*
	Problem: 143. Reorder List
	Reference: https://leetcode.com/problems/reorder-list/
    Complexity: O(n)
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	mapper := make(map[int]*ListNode)
	for i := 0; head != nil; i++ {
		mapper[i] = head
		head = head.Next
	}
	sz := len(mapper)
	for i := 0; i < sz/2; i++ {
		curr := mapper[i]
		next := curr.Next
		last := mapper[sz-1-i]
		curr.Next = last
		last.Next = next
	}
	mapper[sz/2].Next = nil
}