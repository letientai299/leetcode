#
# @lc app=leetcode id=69 lang=python3
#
# [69] Sqrt(x)
#
# https://leetcode.com/problems/sqrtx/description/
#
# algorithms
# Easy (31.04%)
# Total Accepted:    352.4K
# Total Submissions: 1.1M
# Testcase Example:  '4'
#
# Implement int sqrt(int x).
# 
# Compute and return the square root of x, where x is guaranteed to be a
# non-negative integer.
# 
# Since the return type is an integer, the decimal digits are truncated and
# only the integer part of the result is returned.
# 
# Example 1:
# 
# 
# Input: 4
# Output: 2
# 
# 
# Example 2:
# 
# 
# Input: 8
# Output: 2
# Explanation: The square root of 8 is 2.82842..., and since 
# the decimal part is truncated, 2 is returned.
# 
#
#
import math


class Solution:
    def mySqrt(self, x: int) -> int:
        return math.floor(math.sqrt(x))


if __name__ == '__main__':
    tests = [
        (0, 0),
        (1, 1),
        (2, 1),
        (3, 1),
        (4, 2),
        (5, 2),
        (6, 2),
        (7, 2),
        (8, 2),
        (9, 3),
    ]
    for tc in tests:
        actual = Solution().mySqrt(tc[0])
        expected = tc[1]
        if actual != expected:
            print(f"mySqrt({tc[0]}) == {actual} != {expected}")
