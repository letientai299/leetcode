#
# @lc app=leetcode id=83 lang=python3
#
# [83] Remove Duplicates from Sorted List
#
# https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
#
# algorithms
# Easy (42.26%)
# Total Accepted:    317.2K
# Total Submissions: 750.1K
# Testcase Example:  '[1,1,2]'
#
# Given a sorted linked list, delete all duplicates such that each element
# appear only once.
# 
# Example 1:
# 
# 
# Input: 1->1->2
# Output: 1->2
# 
# 
# Example 2:
# 
# 
# Input: 1->1->2->3->3
# Output: 1->2->3
# 
# 
#
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
from list_node import ListNode


class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        if head is None:
            return None

        node = head
        while node.next is not None:
            if node.val != node.next.val:
                node = node.next
                continue

            node.next = node.next.next

        return head


if __name__ == '__main__':
    tests = [
        (ListNode(1, 2, 2, 3), ListNode(1, 2, 3)),
        (ListNode(2, 2, 2, 3), ListNode(2, 3)),
        (ListNode(2, 2, 2, 2), ListNode(2)),
    ]
    for tc in tests:
        actual = Solution().deleteDuplicates(tc[0])
        expected = tc[1]
        if actual != expected:
            print(f"deleteDuplicates({tc[0]}) == {actual} != {expected}")
