#
# @lc app=leetcode id=27 lang=python3
#
# [27] Remove Element
#
# https://leetcode.com/problems/remove-element/description/
#
# algorithms
# Easy (43.59%)
# Total Accepted:    377K
# Total Submissions: 862.9K
# Testcase Example:  '[3,2,2,3]\n3'
#
# Given an array nums and a value val, remove all instances of that value
# in-place and return the new length.
# 
# Do not allocate extra space for another array, you must do this by modifying
# the input array in-place with O(1) extra memory.
# 
# The order of elements can be changed. It doesn't matter what you leave beyond
# the new length.
# 
# Example 1:
# 
# 
# Given nums = [3,2,2,3], val = 3,
# 
# Your function should return length = 2, with the first two elements of nums
# being 2.
# 
# It doesn't matter what you leave beyond the returned length.
# 
# 
# Example 2:
# 
# 
# Given nums = [0,1,2,2,3,0,4,2], val = 2,
# 
# Your function should return length = 5, with the first five elements of nums
# containing 0, 1, 3, 0, and 4.
# 
# Note that the order of those five elements can be arbitrary.
# 
# It doesn't matter what values are set beyond the returned length.
# 
# Clarification:
# 
# Confused why the returned value is an integer but your answer is an array?
# 
# Note that the input array is passed in by reference, which means modification
# to the input array will be known to the caller as well.
# 
# Internally you can think of this:
# 
# 
# // nums is passed in by reference. (i.e., without making a copy)
# int len = removeElement(nums, val);
# 
# // any modification to nums in your function would be known by the caller.
# // using the length returned by your function, it prints the first len
# elements.
# for (int i = 0; i < len; i++) {
# print(nums[i]);
# }
# 
#
from typing import List


class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        current = 0
        runner = 0
        while runner < len(nums):
            if nums[runner] != val:
                nums[current] = nums[runner]
                current += 1

            runner += 1

        return current


if __name__ == '__main__':
    tests = [
        ([1], 0, 1, [1]),
        ([1], 1, 0, []),
        ([], 1, 0, []),
        ([1, 2, 1], 1, 1, [2]),
        ([2, 3, 3, 2], 2, 2, [3, 3]),
        ([0, 1, 2, 2, 3, 0, 4, 2], 2, 5, [0, 1, 3, 0, 4])
    ]
    for i, tc in enumerate(tests):
        arr = tc[0].copy()
        actual = Solution().removeElement(tc[0], tc[1])
        expected = tc[2]
        if actual != expected or tc[0][:actual] != tc[3]:
            print(f"{i}/ removeElement({arr, tc[1]}) == {actual} != {expected}")
            print(f'{i}/ after: {tc[0][:actual]}')
