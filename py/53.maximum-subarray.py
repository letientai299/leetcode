#
# @lc app=leetcode id=53 lang=python3
#
# [53] Maximum Subarray
#
# https://leetcode.com/problems/maximum-subarray/description/
#
# algorithms
# Easy (42.89%)
# Total Accepted:    473.6K
# Total Submissions: 1.1M
# Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
#
# Given an integer array nums, find the contiguous subarray (containing at
# least one number) which has the largest sum and return its sum.
# 
# Example:
# 
# 
# Input: [-2,1,-3,4,-1,2,1,-5,4],
# Output: 6
# Explanation: [4,-1,2,1] has the largest sum = 6.
# 
# 
# Follow up:
# 
# If you have figured out the O(n) solution, try coding another solution using
# the divide and conquer approach, which is more subtle.
# 
#
from typing import List


class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        if not nums:
            return 0

        cur = nums[0]
        maxSub = nums[0]
        for x in nums[1:]:
            if cur > 0:
                cur += x
            else:
                cur = x

            if maxSub < cur:
                maxSub = cur

        return maxSub


if __name__ == '__main__':
    tests = [
        ([], 0),
        ([1], 1),
        ([1, 2], 3),
        ([1, -2], 1),
        ([4, -1, 2, 1], 6),
        ([-2, 1, -3, 4, -1, 2, 1, -5, 4], 6)
    ]
    for tc in tests:
        actual = Solution().maxSubArray(tc[0])
        expected = tc[1]
        if actual != expected:
            print(f"maxSubArray({tc[0]}) == {actual} != {expected}")
