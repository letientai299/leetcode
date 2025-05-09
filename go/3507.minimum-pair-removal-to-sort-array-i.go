package main

import (
  "math"
)

// NOTE: this problem doesn't feel easy, because its accepted time complexity is
// quite high. We can continue to improve the time complexity by using max
//-heap, however, the complete solution would be very complicated.

// Minimum Pair Removal to Sort Array I
//
// Easy
//
// https://leetcode.com/problems/minimum-pair-removal-to-sort-array-i
//
// Given an array `nums`, you can perform the following operation any number of
// times:
//
// - Select the **adjacent** pair with the **minimum** sum in `nums`. If
// multiple such pairs exist, choose the leftmost one.
// - Replace the pair with their sum.
//
// Return the **minimum number of operations** needed to make the array
// **non-decreasing**.
//
// An array is said to be **non-decreasing** if each element is greater than or
// equal to its previous element (if it exists).
//
// **Example 1:**
//
// **Input:** nums = [5,2,3,1]
//
// **Output:** 2
//
// **Explanation:**
//
// - The pair `(3,1)` has the minimum sum of 4. After replacement, `nums =
// [5,2,4]`.
// - The pair `(2,4)` has the minimum sum of 6. After replacement, `nums =
// [5,6]`.
//
// The array `nums` became non-decreasing in two operations.
//
// **Example 2:**
//
// **Input:** nums = [1,2,2]
//
// **Output:** 0
//
// **Explanation:**
//
// The array `nums` is already sorted.
//
// **Constraints:**
//
// - `1 <= nums.length <= 50`
// - `-1000 <= nums[i] <= 1000`
func minimumPairRemoval(nums []int) int {
  if len(nums) <= 1 {
    return 0 // nothing to do
  }

  // looping through the array, do 2 things:
  // - convert array to doubly linked list, so that removing element is faster
  // - compute the heap of adjacent sum
  pre := &ListNode{}
  root := pre
  for _, v := range nums {
    pre.Next = &ListNode{Val: v}
    pre = pre.Next
  }

  // pre is now pointing to the very last node. To avoid nil checking later on,
  // we add a dummy node to the end of the list.
  pre.Next = &ListNode{
    // this is greater than max possible list sum given problem constraints
    Val: 1000*50 + 1,
  }
  root = root.Next

  // helper to check if the list is sorted
  var isSorted func(node *ListNode) bool
  isSorted = func(root *ListNode) bool {
    return root == nil || root.Next == nil ||
      (root.Val <= root.Next.Val && isSorted(root.Next))
  }

  // keep popping the heap, updating the list and stop when the list is sorted
  cnt := 0
  for !isSorted(root) {
    cnt++

    // find the min sum pair
    minVal := math.MaxInt32
    cur := root
    minNode := cur

    for cur.Next != nil {
      v := cur.Val + cur.Next.Val
      if v < minVal {
        minNode = cur
        minVal = v
      }
      cur = cur.Next
    }

    minNode.Val = minVal
    minNode.Next = minNode.Next.Next
  }

  return cnt
}
