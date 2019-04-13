#
# @lc app=leetcode id=7 lang=python3
#
# [7] Reverse Integer
#
# https://leetcode.com/problems/reverse-integer/description/
#
# algorithms
# Easy (25.13%)
# Total Accepted:    623.3K
# Total Submissions: 2.5M
# Testcase Example:  '123'
#
# Given a 32-bit signed integer, reverse digits of an integer.
# 
# Example 1:
# 
# 
# Input: 123
# Output: 321
# 
# 
# Example 2:
# 
# 
# Input: -123
# Output: -321
# 
# 
# Example 3:
# 
# 
# Input: 120
# Output: 21
# 
# 
# Note:
# Assume we are dealing with an environment which could only store integers
# within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of
# this problem, assume that your function returns 0 when the reversed integer
# overflows.
# 
#


class Solution(object):
    @staticmethod
    def reverse(x):
        negative = False
        if x < 0:
            negative = True
            x = -x

        res = int(''.join(reversed(str(x))))
        if negative:
            res = -res

        if res > 2147483647 or res < -2147483647:
            return 0
        else:
            return res


if __name__ == '__main__':
    tests = {
        2 ** 31 + 1: 0,
        -1: -1,
        1: 1,
        0: 0,
        10: 1,
        1000000: 1,
        134: 431
    }
    for param, expected in tests.items():
        actual = Solution().reverse(param)
        if actual != expected:
            print(f"reverse({param}) == {actual} != {expected}")
