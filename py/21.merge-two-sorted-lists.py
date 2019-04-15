#
# @lc app=leetcode id=21 lang=python3
#
# [21] Merge Two Sorted Lists
#
# https://leetcode.com/problems/merge-two-sorted-lists/description/
#
# algorithms
# Easy (45.87%)
# Total Accepted:    521.2K
# Total Submissions: 1.1M
# Testcase Example:  '[1,2,4]\n[1,3,4]'
#
# Merge two sorted linked lists and return it as a new list. The new list
# should be made by splicing together the nodes of the first two lists.
#
# Example:
#
# Input: 1->2->4, 1->3->4
# Output: 1->1->2->3->4->4
#
#
#
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


import os

# Check if we are running locally, and the file existed,
# so that we can import and use it.
if os.path.isfile('./list_node.py'):
    # On leetcode, our file is not existed,
    # thus, this code won't run, # below code will use leetcode provided version
    from list_node import ListNode


# This setup allow our code to run fine unchanged both on leetcode and locally.
# But, due to the IO call (check for file), our code will run slower, and the
# runtime ratio won't be correct.
# So, when submit the final code to leetcode, we must manually remove all the
# checking.


class Solution:
    def mergeTwoLists(self, list1: ListNode, list2: ListNode) -> ListNode:
        head = res = ListNode(0)
        while list2 and list1:
            if list2.val <= list1.val:
                res.next = list2
                list2 = list2.next
            else:
                res.next = list1
                list1 = list1.next

            res = res.next

        while list1:
            res.next = list1
            list1 = list1.next
            res = res.next

        while list2:
            res.next = list2
            list2 = list2.next
            res = res.next

        return head.next


if __name__ == "__main__" and os.path.isfile('./list_node.py'):
    tests = [
        (ListNode(1, 2, 3), None, ListNode(1, 2, 3)),
        (ListNode(1, 2, 3), ListNode(1, 2), ListNode(1, 1, 2, 2, 3)),
        (ListNode(1, 2, 3), ListNode(4, 5), ListNode(1, 2, 3, 4, 5)),
        (ListNode(1, 2, 3), ListNode(1, 5), ListNode(1, 1, 2, 3, 5))
    ]
    for test in tests:
        a = test[0]
        b = test[1]
        expected = test[2]
        actual = Solution().mergeTwoLists(a, b)
        if actual != expected:
            print(f"mergeTwoLists({a}, {b}) is {actual} != {expected}")
