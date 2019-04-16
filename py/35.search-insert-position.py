#
# @lc app=leetcode id=35 lang=python3
#
# [35] Search Insert Position
#
# https://leetcode.com/problems/search-insert-position/description/
#
# algorithms
# Easy (40.66%)
# Total Accepted:    383.7K
# Total Submissions: 943.6K
# Testcase Example:  '[1,3,5,6]\n5'
#
# Given a sorted array and a target value, return the index if the target is
# found. If not, return the index where it would be if it were inserted in
# order.
# 
# You may assume no duplicates in the array.
# 
# Example 1:
# 
# 
# Input: [1,3,5,6], 5
# Output: 2
# 
# 
# Example 2:
# 
# 
# Input: [1,3,5,6], 2
# Output: 1
# 
# 
# Example 3:
# 
# 
# Input: [1,3,5,6], 7
# Output: 4
# 
# 
# Example 4:
# 
# 
# Input: [1,3,5,6], 0
# Output: 0
# 
# 
#
from typing import List


class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        left = 0
        right = len(nums) - 1
        while left < right:
            mid = left + (right - left) // 2

            if nums[mid] == target:
                return mid

            # no duplicate, so this is safe to return
            if nums[mid] < target < nums[mid + 1]:
                return mid + 1

            if nums[mid] > target:
                right = mid
            else:
                left = mid + 1

        if nums[left] < target:
            return left + 1

        return left


if __name__ == '__main__':
    tests = [
        ([1], 0, 0),
        ([1, 4], 3, 1),
        ([1, 2, 4], 3, 2),
        ([1, 2, 4], 2, 1),
        ([1, 2, 4], 1, 0),
        ([1, 2, 4], 0, 0),
        ([1, 2, 4], 5, 3),
        ([1, 2, 3, 4], 3, 2),
    ]
    for tc in tests:
        actual = Solution().searchInsert(tc[0], tc[1])
        expected = tc[2]
        if actual != expected:
            print(f"searchInsert({tc[0], tc[1]}) == {actual} != {expected}")
