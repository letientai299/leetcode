#
# @lc app=leetcode id=67 lang=python3
#
# [67] Add Binary
#
# https://leetcode.com/problems/add-binary/description/
#
# algorithms
# Easy (38.08%)
# Total Accepted:    280.7K
# Total Submissions: 736.3K
# Testcase Example:  '"11"\n"1"'
#
# Given two binary strings, return their sum (also a binary string).
# 
# The input strings are both non-empty and contains only characters 1 or 0.
# 
# Example 1:
# 
# 
# Input: a = "11", b = "1"
# Output: "100"
# 
# Example 2:
# 
# 
# Input: a = "1010", b = "1011"
# Output: "10101"
# 
#


class Solution:
    def addBinary(self, a: str, b: str) -> str:
        return bin(int(a, 2) + int(b, 2))[2:]

    # def addBinary(self, a: str, b: str) -> str:
    #     n = min(len(a), len(b))
    #     carry = 0
    #     i = 1
    #     c = ''
    #     while i <= n:
    #         x = int(a[len(a) - i]) + int(b[len(b) - i]) + carry
    #         carry = x // 2
    #         x %= 2
    #         c += str(x)
    #         i += 1
    #
    #     if n == len(a):
    #         rest = b
    #     else:
    #         rest = a
    #
    #     for i in range(n, len(rest)):
    #         x = carry + int(rest[len(rest) - i - 1])
    #         carry = x // 2
    #         x %= 2
    #         c += str(x)
    #
    #     if carry == 1:
    #         c += str(carry)
    #
    #     return c[::-1]


if __name__ == '__main__':
    tests = [
        ('0', '10', '10'),
        ('1', '10', '11'),
        ('11', '11', '110'),
        ('11', '10', '101'),
    ]
    for tc in tests:
        a = tc[0]
        b = tc[1]
        actual = Solution().addBinary(a, b)
        expected = tc[2]
        if actual != expected:
            print(f"({a}, {b}) == {actual} != {expected}")
