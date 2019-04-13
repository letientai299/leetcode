#
# @lc app=leetcode id=9 lang=python3
#
# [9] Palindrome Number
#
# https://leetcode.com/problems/palindrome-number/description/
#
# algorithms
# Easy (42.00%)
# Total Accepted:    521.4K
# Total Submissions: 1.2M
# Testcase Example:  '121'
#
# Determine whether an integer is a palindrome. An integer is a palindrome when
# it reads the same backward as forward.
# 
# Example 1:
# 
# 
# Input: 121
# Output: true
# 
# 
# Example 2:
# 
# 
# Input: -121
# Output: false
# Explanation: From left to right, it reads -121. From right to left, it
# becomes 121-. Therefore it is not a palindrome.
# 
# 
# Example 3:
# 
# 
# Input: 10
# Output: false
# Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
# 
# 
# Follow up:
# 
# Coud you solve it without converting the integer to a string?
# 
#


class Solution:
    def isPalindrome(self, x: int) -> object:
        if x < 0:
            return False
        # Don't know why but reverse string and convert back to int yield the
        # best performance.
        return int(str(x)[::-1]) - x == 0


if __name__ == '__main__':
    tests = {
        19: False,
        10: False,
        -101: False,
        0: True,
        11: True,
        1111: True,
        101: True,
    }
    for param, expected in tests.items():
        actual = Solution().isPalindrome(param)
        if actual != expected:
            print(f"isPalindrome({param}) == {actual} != {expected}")
