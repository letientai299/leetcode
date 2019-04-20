#
# @lc app=leetcode id=66 lang=python3
#
# [66] Plus One
#
# https://leetcode.com/problems/plus-one/description/
#
# algorithms
# Easy (40.74%)
# Total Accepted:    360.3K
# Total Submissions: 883.9K
# Testcase Example:  '[1,2,3]'
#
# Given a non-empty array of digits representing a non-negative integer, plus
# one to the integer.
# 
# The digits are stored such that the most significant digit is at the head of
# the list, and each element in the array contain a single digit.
# 
# You may assume the integer does not contain any leading zero, except the
# number 0 itself.
# 
# Example 1:
# 
# 
# Input: [1,2,3]
# Output: [1,2,4]
# Explanation: The array represents the integer 123.
# 
# 
# Example 2:
# 
# 
# Input: [4,3,2,1]
# Output: [4,3,2,2]
# Explanation: The array represents the integer 4321.
# 
#
from typing import List


class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        carry = 1
        for i in range(len(digits) - 1, -1, -1):
            x = digits[i]
            x += carry
            carry = x // 10
            x = x % 10
            digits[i] = x

        if carry > 0:
            return [carry] + digits

        return digits


if __name__ == '__main__':
    tests = [
        ([0], [1]),
        ([1, 0], [1, 1]),
        ([8], [9]),
        ([9], [1, 0]),
    ]
    for tc in tests:
        original = tc[0].copy()
        actual = Solution().plusOne(tc[0])
        expected = tc[1]
        if actual != expected:
            print(f"plug({original}) == {actual} != {expected}")
