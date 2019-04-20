#
# @lc app=leetcode id=88 lang=python3
#
# [88] Merge Sorted Array
#
# https://leetcode.com/problems/merge-sorted-array/description/
#
# algorithms
# Easy (35.37%)
# Total Accepted:    350.6K
# Total Submissions: 991.1K
# Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
#
# Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as
# one sorted array.
# 
# Note:
# 
# 
# The number of elements initialized in nums1 and nums2 are m and n
# respectively.
# You may assume that nums1 has enough space (size that is greater or equal to
# m + n) to hold additional elements from nums2.
# 
# 
# Example:
# 
# 
# Input:
# nums1 = [1,2,3,0,0,0], m = 3
# nums2 = [2,5,6],       n = 3
# 
# Output: [1,2,2,3,5,6]
# 
# 
#
from typing import List


class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        # we will merge from the end of the final array, thus, start as the end
        i = len(nums1) - 1

        # for each numbers in nums1 and nums2, pick the smaller, and put it
        # into the end of the final array
        m -= 1
        n -= 1
        while m >= 0 and n >= 0:
            a = nums1[m]
            b = nums2[n]
            if a > b:
                nums1[i] = a
                m -= 1
            else:
                nums1[i] = b
                n -= 1

            i -= 1

        while m >= 0:
            nums1[i] = nums1[m]
            m -= 1
            i -= 1

        while n >= 0:
            nums1[i] = nums2[n]
            n -= 1
            i -= 1

        return


if __name__ == '__main__':
    tests = [
        ([1, 2, 3, 0, 0, 0], 3, [2, 5, 6], 3),
        ([0], 0, [1], 1),
        ([1, 2, 3, 0, 0, 0], 3, [1, 5, 6], 3),
        ([1, 2, 8, 0, 0, 0], 3, [1, 5, 6], 3),
        ([0, 0], 0, [1, 1], 2),
    ]
    for tc in tests:
        nums1 = tc[0]
        nums2 = tc[2]
        Solution().merge(nums1, tc[1], nums2, tc[3])
        for i in range(len(nums1) - 1):
            if nums1[i] > nums1[i + 1]:
                print(f"Array is not merged correctly: {nums1}")
