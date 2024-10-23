/*
	Problem: 206. Reverse Linked List
	Reference: https://leetcode.com/problems/reverse-linked-list/
    Complexity: O(n)
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func buildStream(head *ListNode) []*ListNode {
	stream := []*ListNode{}
	curr := head
	for curr != nil {
		stream = append(stream, curr)
		curr = curr.Next
	}
	return stream
}

func reverseLink(stream []*ListNode) *ListNode {
	if len(stream) == 0 {
		return nil
	}
	head := stream[len(stream)-1]
	curr := head
	for i := len(stream) - 2; i >= 0; i-- {
		curr.Next = stream[i]
		curr = curr.Next
		curr.Next = nil
	}
	return head
}

func reverseList(head *ListNode) *ListNode {
	stream := buildStream(head)
	return reverseLink(stream)
}