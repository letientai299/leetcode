package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNodes(head *ListNode, m int, n int) *ListNode {
	node := &ListNode{Next: head}
	for node != nil {
		for i := 0; i < m && node != nil; i++ {
			node = node.Next
		}

		for i := 0; i < n && node != nil && node.Next != nil; i++ {
			node.Next = node.Next.Next
		}
	}
	return head
}
