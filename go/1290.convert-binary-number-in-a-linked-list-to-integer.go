package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	n := 0
	for head != nil {
		n = n<<1 + head.Val
		head = head.Next
	}
	return n
}
